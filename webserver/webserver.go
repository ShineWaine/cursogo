package webserver

import "net/http"

func MiWebServer() {
	//http.HandleFunc es un manejador de direcciónes.
	//esto maneja las urls. En este caso si lo que se le manda es '/' raiz,
	//enviará la página de inicio.html desde la función de 'home'
	http.HandleFunc("/", home)        //define la ruta '/'
	http.ListenAndServe(":3000", nil) //define puerto de escucha
}

/*
Cuando una función tiene que procesar una peticion de web, debe procesar
dos párámetros obligatorios:
  - http.ResponseWriter (writer)
  - *http.Request  (request) *-> es un puntero a una dirección de memoria
*/
func home(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile, lo que hace es servir la página web usando los
	//parametros obtenidos como entrada en la llamada a la función.
	http.ServeFile(w, r, "./webserver/index.html")
}
