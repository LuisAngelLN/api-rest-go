package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type especialidaditl struct {
	ID           string `json:"id"`
	Especialidad string `json:"especialidad"`
	Carrera      string `json:"carrera"`
	Materias     string `json:"materias"`
}

var especialidades = []especialidaditl{
	{ID: "ISIE-DAE-2020-02", Especialidad: "Desarrollo de Aplicaciones Empresariales", Carrera: "Ingeniería en Sistemas", Materias: "Base de Datos Avanzadas, Ciencia de datos, Cómputo y servicios en la nube, Desarrollo de Aplicaciones Móviles, Desarrollo de Aplicaciones Empresariales, Lab. Despliegue de Aplicaciones"},
	{ID: "ISIE-RAE-2020-01", Especialidad: "Redes Convergentes de Alta Disponibilidad con Tecnologías Emergentes", Carrera: "Ingeniería en Sistemas", Materias: "Diseño de Redes, Redes Inalámbricas, Seguridad en Redes, Infraestructura para despliegue de aplicaciones, Fundamentos de IoT"},
	{ID: "IINE-IMC-2020-01", Especialidad: "Manufactura y Calidad", Carrera: "Ingeniería Industrial", Materias: "Sistemas Neumáticos e Hidráulicos, Diseño asistido por computadora, Controladores Lógicos Programables, Robótica Industrial, Medición de la Productividad, Core Tool, Temas Selesctos de Manufactura, Manufactura Flexible"},
	{ID: "IINE-MAP-2020-02", Especialidad: "Manufactura en Artículos de Piel", Carrera: "Ingeniería Industrial", Materias: "Diseño y Modelado, Diseño Asistido por Computadora, Tecnología y Taller, Temas Selectos de Manufactura, Core Tool, Tecnología y Taller II, Medición y Mejoramiento de la Productividad, Administración de Sistemas de Calzado"},
	{ID: "ITIE-GST-2020-01", Especialidad: "Gestión de Servicios de T.I. en Ambientes Empresariales", Carrera: "TICS", Materias: "Seguridad Informática, Administración de Servidores, Gestión de Sistemas VoIP, Estrategias de Gestión de Servicios de TI, Virtualización & IoT"},
	{ID: "GEE-CIP-2017-01", Especialidad: "Gestión de la calidad e innovación de procesos", Carrera: "Ingeniería en Gestión Empresarial", Materias: "Estrategias Fin. y Costos de Calidad, Innovación de Procesos, Calidad Aplicada a la Gestión Empresarial II, Seminario de Calidad, Innovación de Procesos, Seminario de Consultoria, Gestión del Conocimiento, Sem. Gestión Estratégica"},
}

func getEspecialidades(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, especialidades)
}

func postEspecialidades(c *gin.Context) {
	var nuevaEspecialidad especialidaditl

	if err := c.BindJSON(&nuevaEspecialidad); err != nil {
		return
	}

	especialidades = append(especialidades, nuevaEspecialidad)

	c.IndentedJSON(http.StatusCreated, especialidades)

}

func getByID(c *gin.Context) {
	id := c.Param("id")

	for _, e := range especialidades {
		if e.ID == id {
			c.IndentedJSON(http.StatusOK, e)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Especialidad no encontrada."})
}

func main() {
	router := gin.Default()
	router.GET("/especialidades", getEspecialidades)
	router.POST("/especialidades", postEspecialidades)
	router.GET("/especialidades/:id", getByID)

	router.Run("0.0.0.0:8080")
}
