package router

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"html/template"
	"picbed-go/logger"
	"picbed-go/server/controller"
	"picbed-go/server/middleware"
	"picbed-go/setting"
	"picbed-go/static/bindata"
	"strings"
)

func Setup() *gin.Engine {
	// 创建服务
	r := gin.Default()
	r.MaxMultipartMemory = 20 << 20 // 8 MiB

	if setting.App.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	{
		// 配置日志
		r.Use(logger.GinLogger(), logger.GinRecovery(true))
	}

	// 全局异常处理
	r.Use(middleware.Recover)

	// 前端静态文件
	{
		// 加载模板文件
		t, err := loadTemplate()
		if err != nil {
			panic(err)
		}
		r.SetHTMLTemplate(t)

		// 加载静态文件
		fs := assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: nil,
			Prefix:    "assets",
		}
		r.StaticFS("/static", &fs)

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	{
		op := r.Group("/api")
		op.GET("/info", controller.InfoHandler)
		op.POST("/upload/:id", controller.UploadHandler)
	}

	return r
}

// loadTemplate 加载模板文件
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := bindata.Asset(name)
		if err != nil {
			continue
		}
		name := strings.Replace(name, "assets/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
