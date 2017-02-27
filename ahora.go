package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var cloud map[string]string = make(map[string]string)

func encoderStatNow(w http.ResponseWriter, r *http.Request) {

	cookie, err3 := r.Cookie(CookieName)
	if err3 != nil {
		return
	}
	key := cookie.Value
	usr, ok := user[key] // De aquí podemos recoger el usuario
	if !ok {
		return
	}
	username := usr
	anio, mes, dia := time.Now().Date()
	fecha := fmt.Sprintf("%02d/%02d/%02d", dia, mes, anio)
	hh, mm, _ := time.Now().Clock()
	hora := fmt.Sprintf("%02d:%02d", hh, mm)
	tiempo_limite := time.Now().Unix() - 6 //tiempo limite de 6 seg
	db_mu.RLock()
	query, err := db.Query("SELECT streamname, isocode, ip, country, time, bitrate, info FROM encoders WHERE username = ? AND timestamp > ?", username, tiempo_limite)
	db_mu.RUnlock()
	if err != nil {
		Error.Println(err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><p><b>Conectados el día %s a las %s UTC</b></p><table class=\"table table-striped table-bordered table-hover\"><th>Play</th><th>INFO</th><th>País</th><th>IP</th><th>Stream</th><th>Tiempo conectado</th>", username, fecha, hora)
	for query.Next() {
		var isocode, country, streamname, ip, time_connect, info string
		var tiempo, bitrate int
		err = query.Scan(&streamname, &isocode, &ip, &country, &tiempo, &bitrate, &info)
		if err != nil {
			Warning.Println(err)
		}
		isocode = strings.ToLower(isocode)
		time_connect = secs2time(tiempo)
		INFO := fmt.Sprintf("%s [%d kbps]", info, bitrate/1000)
		fmt.Fprintf(w, "<tr><td><a href=\"javascript:launchRemote('play.cgi?stream=%s')\"><img src='images/play.jpg' border='0' title='Play %s'/></a></td><td>%s</td><td><img src=\"images/flags/%s.png\" title=\"%s\"></td><td>%s</td><td>%s</td><td>%s</td></tr>",
			streamname, streamname, INFO, isocode, country, ip, streamname, time_connect)
	}
	query.Close()
	fmt.Fprintf(w, "</table>")
}

func playerStatNow(w http.ResponseWriter, r *http.Request) {
	cookie, err3 := r.Cookie(CookieName)
	if err3 != nil {
		return
	}
	key := cookie.Value
	usr, ok := user[key] // De aquí podemos recoger el usuario
	if !ok {
		return
	}
	username := usr
	var contador int
	tiempo_limite := time.Now().Unix() - 30 //tiempo limite de 30 seg
	db_mu.RLock()
	err := db.QueryRow("SELECT count(*) FROM players WHERE username = ? AND timestamp > ? AND time > 0", username, tiempo_limite).Scan(&contador)
	db_mu.RUnlock()
	if err != nil {
		Error.Println(err)
	}
	if contador >= 100 {
		db_mu.RLock()
		query, err := db.Query("SELECT isocode, country, count(ipclient), streamname FROM players WHERE username = ? AND timestamp > ? AND time > 0 GROUP BY isocode, streamname", username, tiempo_limite)
		db_mu.RUnlock()
		if err != nil {
			Error.Println(err)
		}
		fmt.Fprintf(w, "<table class=\"table table-striped table-bordered table-hover\"><th>País</th><th>Cantidad de IPs</th><th>Stream</th>")
		for query.Next() {
			var isocode, country, ips, streamname string
			err = query.Scan(&isocode, &country, &ips, &streamname)
			if err != nil {
				Warning.Println(err)
			}
			fmt.Fprintf(w, "<tr><td>%s <img class='pull-right' src=\"images/flags/%s.png\" title=\"%s\"></td><td>%s</td><td>%s</td></tr>",
				country, isocode, country, ips, streamname)
		}
		query.Close()
		fmt.Fprintf(w, "<tr><td align=\"center\" colspan='7'><b>Total:</b> %d players conectados</td></tr></table>", contador)
	} else {
		db_mu.RLock()
		query, err := db.Query("SELECT isocode, country, region, city, ipclient, os, streamname, time FROM players WHERE username = ? AND timestamp > ? AND time > 0", username, tiempo_limite)
		db_mu.RUnlock()
		if err != nil {
			Warning.Println(err)
		}
		fmt.Fprintf(w, "<table class=\"table table-striped table-bordered table-hover\"><th>País</th><th>Region</th><th>Ciudad</th><th>Dirección IP</th><th>Stream</th><th>O.S</th><th>Tiempo conectado</th>")
		for query.Next() {
			var isocode, country, region, city, ipclient, os, streamname, time_connect string
			var tiempo int
			err = query.Scan(&isocode, &country, &region, &city, &ipclient, &os, &streamname, &tiempo)
			if err != nil {
				Warning.Println(err)
			}

			isocode = strings.ToLower(isocode)
			time_connect = secs2time(tiempo)
			fmt.Fprintf(w, "<tr><td>%s <img class='pull-right' src=\"images/flags/%s.png\" title=\"%s\"></td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>",
				country, isocode, country, region, city, ipclient, streamname, os, time_connect)
		}
		query.Close()
		fmt.Fprintf(w, "<tr><td align=\"center\" colspan='8'><b>Total:</b> %d players conectados</td></tr></table>", contador)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	cookie, err3 := r.Cookie(CookieName)
	if err3 != nil {
		return
	}
	key := cookie.Value
	usr, ok := user[key] // De aquí podemos recoger el usuario
	if !ok {
		return
	}
	username := usr
	loadSettings(playingsRoot)
	r.ParseForm() // recupera campos del form tanto GET como POST
	allname := username + "-" + r.FormValue("stream")
	stream := "http://" + cloud["cloudserver"] + "/live/" + allname + ".m3u8"
	video := fmt.Sprintf("<script type='text/javascript' src='http://www.domainplayers.org/js/jwplayer.js'></script><div id='container'><video width='600' height='409' controls autoplay src='%s'/></div><script type='text/javascript'>jwplayer('container').setup({ width: '600', height: '409', skin: 'http://www.domainplayers.org/newtubedark.zip', plugins: { 'http://www.domainplayers.org/qualitymonitor.swf' : {} }, image: '', modes: [{ type:'flash', src:'http://www.domainplayers.org/player.swf', config: { autostart: 'true', provider:'http://www.domainplayers.org/HLSProvider5.swf', file:'%s' } }]});</script>", stream, stream)
	fmt.Fprintf(w, "%s", video)
}
