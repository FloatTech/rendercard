package rendercard

const (
	// DefaultWidth 默认宽度
	DefaultWidth = 1272.0
)

// Title 标题配置
type Title struct {
	// Line 行数
	Line int
	// IsEnabled 状态
	IsEnabled bool
	// LeftTitle 左侧标题
	LeftTitle string
	// LeftSubtitle 左侧副标题
	LeftSubtitle string
	// RightTitle 右侧标题
	RightTitle string
	// RightSubtitle 右侧副标题
	RightSubtitle string
	// ImagePath 图片路径
	ImagePath string
	// TitleFont 标题字体路径
	TitleFont string
	// TextFont 正文字体路径
	TextFont string
}

type Alignment uint8

const (
	NilAlign Alignment = iota
	AlignLeft
	AlignCenter
	AlignRight
)

// Card 卡片配置
type Card struct {
	// Width 宽度,默认600
	Width int
	// Height 高度,默认由Title+Text内容决定
	Height int
	// BackgroundImage 背景图
	BackgroundImage string
	// TitleFont 标题字体
	TitleFont string
	// TextFont 正文字体
	TextFont string
	// Title 标题内容
	Title string
	// CanTitleShown 是否显示标题
	CanTitleShown bool
	// IsTextSplitPerElement true为每个元素按行显示,false按空格分割显示;
	IsTextSplitPerElement bool
	// TitleAlign 标题布局[Left|Center|Right],默认Left
	TitleAlign Alignment
	// Text 正文内容
	Text []string
}
