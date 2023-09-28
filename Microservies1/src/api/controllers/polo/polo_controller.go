package polo

const (
	polo = "polo"
)

func Polo(c *gin.Context) {
	c.String(http.StatusOk, polo)
}
