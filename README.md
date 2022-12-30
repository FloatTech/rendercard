<div align="center">
  <img src=".github/warma.png" alt="沃玛" width = "300">
  </a><br>
  <h1>RenderCard</h1>
</div>

 ## Drawtitle

**示例代码**

```go
  package main
     
  import (
      "github.com/Coloured-glaze/gg"
      "github.com/FloatTech/rendercard"
  )

  func main(){
      img, err = rendercard.Titleinfo{
          Line:          9,
          Lefttitle:     "服务列表",
          Leftsubtitle:  "service_list",
          Righttitle:    "FloatTech",
          Rightsubtitle: "ZeroBot-Plugin",
          Fontpath:      "/font.ttf",
          Imgpath:       "/kanban.png",
      }.Drawtitle()
      if err != nil {
        panic(err)
      }
      gg.SavePNG("/Drawtitle.png",img)
  }
 ```
 **示例图**

<img src=".github/Drawtitle.png" alt="服务列表" width = "">

## Drawtitledtext

**示例代码**

```go
  package main
     
  import (
      "github.com/Coloured-glaze/gg"
      "github.com/FloatTech/rendercard"
  )

  func main(){
      textlist := []string{"文字描述1", "文字描述2", "文字描述3"}
      img, err = rendercard.Titleinfo{
          Lefttitle:     "名称"
          Leftsubtitle:  "简介"
          Righttitle:    "FloatTech",
          Rightsubtitle: "ZeroBot-Plugin",
          Imgpath:       "/kanban.png",
          Fontpath:      "/font.ttf",
          Fontpath2:     "/font2.ttf",
          Status:        true,
      }.Drawtitledtext(textlist)
      if err != nil {
        panic(err)
      }
      gg.SavePNG("/Drawtitledtext.png",img)
  }
 ```
 **示例图**

<img src=".github/Drawtitledtext.png" alt="" width = "">

## Drawcard

**示例代码**

```go
  package main
     
  import (
      "github.com/Coloured-glaze/gg"
      "github.com/FloatTech/rendercard"
  )

  func main(){
      img, err = rendercard.Titleinfo{
        Lefttitle:    "名称",				
        Leftsubtitle: "简介",
        Imgpath:      "/banner.png",
        Fontpath:     "/font.ttf",
        Fontpath2:    "/font2.ttf",
        Status:       true,
      }.Drawtitledtext()
      if err != nil {
        panic(err)
      }
      gg.SavePNG("/Drawcard.png",img)
  }
 ```
 **示例图**

<img src=".github/Drawcard.png" alt="" width = "">