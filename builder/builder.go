package builder

/*
建造者模式
*/
//房子【产品（Product）】
type House struct {
	Foundation string //地基
	Structure  string //结构
	Roof       string //屋顶
	Interior   string //装饰
}

func (h *House) SetFoundation(foundation string) {
	h.Foundation = foundation
}

func (h *House) SetStructure(structure string) {
	h.Structure = structure
}

func (h *House) SetRoof(roof string) {
	h.Roof = roof
}

func (h *House) SetInterior(interior string) {
	h.Interior = interior
}

//房子抽象构建者【抽象建造者（Abstract Builder）】
type HouseBuilder interface {
	BuildFoundation()
	BuildStructure()
	BuildRoof()
	BuildInterior()
}

//房子具体构建者【具体建造者（Concrete Builder）】
type ConcreteHouseBuilder struct {
	house *House
}

func (builder *ConcreteHouseBuilder) BuildFoundation() {
	builder.house.SetFoundation("Standard Foundation")
}

func (builder *ConcreteHouseBuilder) BuildStructure() {
	builder.house.SetStructure("Standard Structure")
}

func (builder *ConcreteHouseBuilder) BuildRoof() {
	builder.house.SetRoof("Standard Roof")
}

func (builder *ConcreteHouseBuilder) BuildInterior() {
	builder.house.SetInterior("Standard Interior")
}

//豪宅构建者Builder【具体建造者（Concrete Builder）】
type LuxuryHouseBuilder struct {
	house *House
}

func (builder *LuxuryHouseBuilder) BuildFoundation() {
	builder.house.SetFoundation("Strong Foundation")
}

func (builder *LuxuryHouseBuilder) BuildStructure() {
	builder.house.SetStructure("Reinforced Structure")
}

func (builder *LuxuryHouseBuilder) BuildRoof() {
	builder.house.SetRoof("Reinforced Roof")
}

func (builder *LuxuryHouseBuilder) BuildInterior() {
	builder.house.SetInterior("Luxury Interior")
}

//指导者（Director）
type Director struct {
	builder HouseBuilder
}

func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Build() {
	d.builder.BuildFoundation()
	d.builder.BuildStructure()
	d.builder.BuildRoof()
	d.builder.BuildInterior()
}
