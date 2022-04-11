package main

const (
	/*Algunas constantes usadas para respuesta en JSON. Verificar constantes.js*/
	RespuestaErrorNegocioExistente          = 0
	RespuestaErrorRegistrandoNegocio        = 1
	RespuestaNegocioRegistradoCorrectamente = 2
	RespuestaLoginNegocioNoVerificado       = 3
	RespuestaLoginCorrecto                  = 4
	RespuestaLoginError                     = 5
	RespuestaLoginIncorrecto                = 6
	// Constantes del sistema
	ImpresionNativa                      = false
	TamanioMaximoArchivoImportacionExcel = 5 << 20 // 5MB
	NombreArchivoExportadoCSV            = "ProductosExportados_STO.csv"
	NombreArchivoExportadoExcel          = "ProductosExportados_STO.xlsx"
	LimiteProductosMasVendidos           = 10
	DiasParaApartadosProximosAVencer     = 7
	LimiteProductosMenosVendidos         = 10
	LimiteProductosNuncaVendidos         = 10
	LimiteAutoCompletadoClientes         = 10
	MensajePieDeTicket                   = "*** Agradecemos su preferencia ***"
	MIME                                 = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	NumeroAvisosAntesDeEliminarCuenta    = 3
	DiasParaMarcarCuentaComoInactiva     = 7
	AsuntoReactivarCuenta                = "Reactiva tu cuenta de POS-STO"
	MensajeCompartir                     = "Conoce POS - STO. Un punto de venta completo."
	UrlAcortadaCompartirSistema          = "https://github.com/Joss852"
	RutaDirectorioContenidoEstatico      = "/static/"
	NombreSesion                         = "PosSTOSesion"
	NombreBaseDeDatosSesiones            = "./sesiones.sqlite"
	ClaveCookie                          = "905ff3f39809b9f9e204204596e234589a99a8fab5e691c08474584a5f3b7523"
	EdadDeSesionEnSegundos               = 60 * 60 * 24 // Un día en segundos
	DominioPermitidoCORS                 = "http://localhost:8080"
	PuertoServidor                       = ":2106" // Debe coincidir con constantes.js
	RutaGeneralNoNecesitaComprobacion    = "RUTA_NO_AUTH"
	ControladorBD                        = "sqlite3"
	NombreDirectorioParaSubidas          = "subidas"
	CadenaConexionBDUsuarios             = "negocio_jm.me_"
	CadenaConexionBDNegocios             = "negocios_STO_by_jossm.me.db"
	PrefijoBDNegocios                    = "sto"
	NombreArchivoEsquemaSQLInit          = "init.sql"
	NombreArchivoEsquemaSQLNegocios      = "esquema_negocios_sqlite.sql"
	NombreArchivoEsquemaSQLSistemas      = "esquema_spos_sqlite.sql"
	NombrePrimerUsuarioAdmin             = "root"
	ClaveAPICorreo                       = ""
	PrefijoRutasAdmin                    = "/auth"
	GmailServidor                        = ""
	GmailPuerto                          = ""
	GmailCorreo                          = ""
	GmailPass                            = ""
	UrlBaseApp                           = ""
	UrlBasePaginaWeb                     = "https://joss852.github.io/"
	CorreoSoporteYContacto               = "jjosseph.morales@gmail.com"
)
