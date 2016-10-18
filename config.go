package main

const (
	// variables de configuracion del servidor HTTP
	rootdir                = "/var/segments/"              // raiz de nuestro sitio web
	session           bool = true                          // habilitado el control de sesiones por cookies
	session_timeout        = 600                           // segundos de timeout de una session
	first_page             = "index"                       // Sería la página de login (siempre es .html)
	enter_page             = "ahora.html"                  // Sería la página de entrada tras el login
	enter_page_admin       = "monthly_admin.html"          // Sería la página de entrada al panel admin tras el login
	http_port              = "80"                          // puerto del server HTTP
	name_username          = "user"                        // name del input username en la página de login
	name_password          = "password"                    // name del input password en la página de login
	CookieName             = "GOSESSID"                    // nombre del cookie que guardamos en el navegador del usuario
	login_cgi              = "/login.cgi"                  // action cgi login in login page
	logout_cgi             = "/logout.cgi"                 // logout link at any page
	session_value_len      = 26                            // longitud en caracteres del Value de la session cookie
	spanHTMLlogerr         = "<span id='loginerr'></span>" // <span> donde publicar el mensaje de error de login
	ErrorText              = "Error de Login"              // mensaje a mostrar tras un error de login en la pagina de login
	logFile                = "/var/log/hlserver.log"       //ruta del archivo de errores
	DirDB                  = "/usr/local/bin/live.db"
	daylyDB                = "/usr/local/bin/dayly.db"         // base de datos de mantenimiento dirario
	monthlyDB              = "/usr/local/bin/monthly.db"       // base de datos de mantenimiento mensual
	dirDaylys              = "/usr/local/bin/daylys/"          // directorio donde se van almacenar las BDs de mantenimiento diario
	dirMonthlys            = "/usr/local/bin/monthlys/"        // directorio donde se van almacenar las BDs de mantenimiento mensual
	playingsRoot           = "/usr/local/bin/playings.reg"     // fichero que contiene el nombre del cloud
	dirGeoip               = "/usr/local/bin/GeoIP2-City.mmdb" // directorio donde se carga el fichero de geolocalización
)
