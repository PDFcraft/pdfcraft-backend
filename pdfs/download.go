package pdfs

import (
	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileid := c.Param("fileid")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
	c.Next()
	originFileName := db.GetFileNamePair(fileid + ".pdf")
	c.FileAttachment("./files/output/"+fileid+".pdf", originFileName)

}
