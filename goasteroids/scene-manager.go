package goasteroids

import "github.com/hajimehoshi/ebiten/v2"

var (
	//离场幕布
	transitionFrom = ebiten.NewImage(ScreenWidth, ScreenHeight) //旧幕布
	transitionTo   = ebiten.NewImage(ScreenWidth, ScreenHeight) //新幕布
)

var transitionMaxCount int = 25 //最大过渡步长

//每个场景必须满足ebiten rungame接口
type Scene interface {
	Update(state *State) error
	Draw(screen *ebiten.Image)
}

//状态管理
type State struct {
	SceneManager *SceneManager
	Input        *Input
}

//场景管理
type SceneManager struct {
	current         Scene //当前场景
	next            Scene //下一个场景
	transitionCount int   //当前过渡步长
}

//绘制场景
func (s *SceneManager) Draw(r *ebiten.Image) {
	//如果当前没有在过渡，则绘制当前场景
	if s.transitionCount == 0 {
		s.current.Draw(r)
		return
	}
	//else
	transitionFrom.Clear()         //清空窗口内容
	s.current.Draw(transitionFrom) //将当前场景绘制到幕布transitionFrom

	transitionTo.Clear()
	s.next.Draw(transitionTo) //将下一个场景绘制到幕布transitionTo

	//先绘制的在最底
	r.DrawImage(transitionFrom, nil)

	//新场景慢慢显现
	alpha := 1 - float32(s.transitionCount)/float32(transitionMaxCount) //透明度 | 0->1 逐渐不透明
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)

	//幕布transitionTo 显示在最顶
	r.DrawImage(transitionTo, op)

}

//更新场景
func (s *SceneManager) Update(_ *Input) error {
	if s.transitionCount == 0 {
		return s.current.Update(&State{
			SceneManager: s,
		})
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}
	//if s.transitionCount <0
	s.current = s.next
	s.next = nil
	return nil

}

//切换场景
func (s *SceneManager) GoToScene(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}

}
