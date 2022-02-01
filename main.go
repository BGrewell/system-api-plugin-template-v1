package main

import (
  "github.com/BGrewell/system-api-plugin"
  "github.com/gin-gonic/gin"
)

func Load() plugin.Plugin {
  return &TemplatePlugin{}
}

type TemplatePlugin struct {
  Message string
}

func (p *TemplatePlugin) Register(config map[string]interface{}, r *gin.RouterGroup) error {
  // Load configuration if it exists
  p.Message = "Template Plugin"
  if config != nil {
    if msg, ok := config["message"]; ok {
      p.Message = msg.(string)
    }
  }
  
  // Register routes
  r.GET("/", p.HomeHandler)
  
  return nil
}

func (p *TemplatePlugin) Name() string {
  return "Template"
}

func (p *TemplatePlugin) HomeHandler(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": p.Message,
  })
}
