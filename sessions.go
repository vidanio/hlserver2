package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// mapas de control de sessions
var user map[string]string = make(map[string]string)
var ip map[string]string = make(map[string]string)
var tiempo map[string]time.Time = make(map[string]time.Time)

func controlinternalsessions() {
	for {
		for k, v := range tiempo {
			if (time.Since(v).Seconds() + float64(session_timeout)) > float64(session_timeout) {
				delete(user, k)
				delete(ip, k)
				delete(tiempo, k)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// genera una session id o Value del Cookie aleatoria y de la longitud que se quiera
func sessionid(r *rand.Rand, n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

// funcion q tramita el login correcto o erroneo
func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() // recupera campos del form tanto GET como POST

	usuario := r.FormValue(name_username)
	pass := r.FormValue(name_password)
	// Autenticar para panel de cliente
	if authentication(usuario, pass) {
		// Generamos la Cookie a escibir en el navegador del usuario
		aleat := rand.New(rand.NewSource(time.Now().UnixNano()))
		sid := sessionid(aleat, session_value_len)
		expiration := time.Now().Add(time.Duration(session_timeout) * time.Second)
		cookie := http.Cookie{Name: CookieName, Value: sid, Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Println("Escribo Cookie")
		// Guardamos constancia de la session en nuestros mapas internos
		user[sid] = usuario
		ip[sid] = getip(r.RemoteAddr)
		tiempo[sid] = expiration
		// Enviamos a la pagina de entrada tras el login correcto
		http.Redirect(w, r, "/"+enter_page, http.StatusFound)
	} else {
		// Autenticar para panel admin
		if authentication_admin(usuario, pass) {
			// Generamos la Cookie a escibir en el navegador del usuario
			aleat := rand.New(rand.NewSource(time.Now().UnixNano()))
			sid := sessionid(aleat, session_value_len)
			expiration := time.Now().Add(time.Duration(session_timeout) * time.Second)
			cookie := http.Cookie{Name: CookieName, Value: sid, Expires: expiration}
			http.SetCookie(w, &cookie)
			fmt.Println("Escribo Cookie")
			// Guardamos constancia de la session en nuestros mapas internos
			user[sid] = usuario
			ip[sid] = getip(r.RemoteAddr)
			tiempo[sid] = expiration
			// Enviamos a la pagina de entrada tras el login correcto
			http.Redirect(w, r, "/"+enter_page_admin, http.StatusFound)
		} else {
			// Te devolvemos a la pagina inicial de login
			fmt.Println("Login incorrecto")
			http.Redirect(w, r, "/"+first_page+".html?err", http.StatusFound)
		}
	}
}

// función que tramita el logout de la session
func logout(w http.ResponseWriter, r *http.Request) {

	cookie, err3 := r.Cookie(CookieName)

	if err3 != nil {
		http.Redirect(w, r, "/"+first_page+".html", http.StatusFound)
	} else {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		delete(user, cookie.Value)
		delete(ip, cookie.Value)
		delete(tiempo, cookie.Value)

		http.Redirect(w, r, "/"+first_page+".html", http.StatusFound)
	}

}

// separa la IPv4/6 del puerto usado con la misma
func getip(pseudoip string) string {
	var res string
	if strings.Contains(pseudoip, "]:") {
		part := strings.Split(pseudoip, "]:")
		res = part[0]
		res = res[1:]
	} else {
		part := strings.Split(pseudoip, ":")
		res = part[0]
	}
	return res
}

// convierte un string numérico en un entero int
func toInt(cant string) (res int) {
	res, _ = strconv.Atoi(cant)
	return
}
