package main

type Player struct {
	currPos Vec2
	targetPos Vec2
	speed float32
}

func (p *Player) MoveTo(v Vec2)  {
	p.targetPos = v
}

func (p *Player) Pos() Vec2  {
	return p.currPos
}

func (p *Player) IsArrived() bool {
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currPos).Normalize()
		newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		//currPos:   Vec2{0,0},
		//targetPos: Vec2{3,6},
		speed:     1,
	}
}

