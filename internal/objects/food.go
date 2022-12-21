package objects

import (
	"math/rand"
	"zeroground/pkg/colors"
	"zeroground/pkg/physics"
	"zeroground/pkg/ui"

	"github.com/veandco/go-sdl2/sdl"
)

type freshness uint8

// food states
const (
	FoodFresh freshness = iota
	FoodStale
	FoodRotten
	FoodInvalid
)

var (
	foodColor = map[freshness]sdl.Color{
		FoodFresh:  colors.New(10, 147, 150),
		FoodStale:  colors.New(238, 155, 0),
		FoodRotten: colors.New(187, 62, 3),
	}

	foodTiming = map[freshness]uint64{
		FoodFresh:  3000,
		FoodStale:  7000,
		FoodRotten: 12000,
	}

	foodSize      = int32(16)
	foodSpawnTime = uint64(5000)
)

type (
	food struct {
		center    sdl.Point
		state     freshness
		createdAt uint64
	}

	FoodSpawner struct {
		lastSpawned uint64
		foods       []*food
	}
)

func (f *food) State() freshness {
	return f.state
}

func (f *food) Position() sdl.Point {
	return f.center
}

func (f *food) Rect() []sdl.Rect {
	return []sdl.Rect{
		{
			X: f.center.X - foodSize/2,
			Y: f.center.Y - foodSize/2,
			W: foodSize,
			H: foodSize,
		},
	}
}

func (f *food) Hitbox() []physics.Plane2D {
	p := physics.NewPlane2D(sdl.Rect{
		X: f.center.X - foodSize/2,
		Y: f.center.Y - foodSize/2,
		W: foodSize,
		H: foodSize,
	})
	return []physics.Plane2D{*p}
}

func (f *food) Update() {
	elapsed := sdl.GetTicks64() - f.createdAt
	duration := foodTiming[f.state]
	if duration <= elapsed {
		f.state += 1
	}
}

func (f *food) Draw(renderer *sdl.Renderer) {
	if f.state == FoodInvalid {
		return
	}
	renderer.SetDrawColor(colors.RGBA(foodColor[f.state]))
	renderer.FillRect(&f.Rect()[0])
	renderer.SetDrawColor(233, 216, 166, 0)
	renderer.DrawRect(&f.Rect()[0])
}

func (f *food) Reset() {
	f.center.X = rand.Int31() % (ui.WindowWidth - 2*foodSize)
	f.center.Y = rand.Int31() % (ui.WindowHeight - 2*foodSize)
	f.state = FoodFresh
	f.createdAt = sdl.GetTicks64()
}

// Position returns the center of most recently appened food
func (f *FoodSpawner) Position() sdl.Point {
	return f.foods[len(f.foods)-1].center
}

func (f *FoodSpawner) Rect() []sdl.Rect {
	rects := make([]sdl.Rect, len(f.foods), len(f.foods))
	for _, food := range f.foods {
		if food.state == FoodInvalid {
			continue
		}
		rects = append(rects, food.Rect()...)
	}
	return rects
}

func (f *FoodSpawner) Handle() {}

// removeInvalidFood removes invlaid foods in O(n) time and O(1) space. This is a brilliant.
// Can add spawn logic in this loop too via resetting food?
func (f *FoodSpawner) removeInvalidFood() {
	// remove invalid turned foods
	removed := 0
	for i, food := range f.foods {
		if removed > 0 {
			f.foods[i-removed] = food
		}
		if food.state == FoodInvalid {
			removed += 1
		}
	}
	f.foods = f.foods[:len(f.foods)-removed]
}

// Update will spawn new a food wheneven one goes stale and removes invalid foods
func (f *FoodSpawner) Update() {
	// instead of creating a new slice for removing rotten food we just do 2 loops
	// better solution would be to create a food pool and reset them.
	// 1st loop is here
	spawn := 0
	for _, food := range f.foods {
		prev := food.state
		if prev == FoodInvalid {
			continue
		}
		food.Update()
		// if food just turned stale
		if prev == FoodFresh && food.state == FoodStale {
			spawn += 1
		}
	}
	// 2nd loop is here
	f.removeInvalidFood()
	// add new food(s)
	if spawn < 1 && (sdl.GetTicks64()-f.lastSpawned) > foodSpawnTime {
		spawn = 1
	}
	for spawn > 0 {
		f.spawnFood()
		spawn -= 1
	}
}

func (f *FoodSpawner) spawnFood() {
	food := &food{}
	food.Reset()
	f.foods = append(f.foods, food)
	f.lastSpawned = sdl.GetTicks64()
	// fmt.Println("ss1:", len(f))
}

func (f FoodSpawner) Draw(renderer *sdl.Renderer) {
	// fmt.Printf("len %d\n", len(f))
	for _, food := range f.foods {
		if food.state == FoodInvalid {
			continue
		}
		food.Draw(renderer)
	}
}

func (f *FoodSpawner) Reset() {
	f.foods = f.foods[:0]
	f.spawnFood()
}

func (fs *FoodSpawner) IntersectingFood(p physics.PhysicsObject2D) *food {
	at := -1
	for i, f := range fs.foods {
		if f.state == FoodInvalid {
			continue
		}
		if physics.HasIntersection(f, p) {
			at = i
			break
		}
	}

	if at < 0 {
		return nil
	}

	f := fs.foods[at]
	if f.state == FoodFresh {
		fs.spawnFood()
	}
	// f.state = FoodInvalid
	fs.foods = append(fs.foods[:at], fs.foods[at+1:]...)
	return f
}

func NewFoodSpawner() *FoodSpawner {
	f := &FoodSpawner{}
	f.Reset()
	return f
}
