// Code generated by "core generate -add-types"; DO NOT EDIT.

package physics

import (
	"cogentcore.org/core/math32"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.BBox", IDName: "b-box", Doc: "BBox contains bounding box and other gross object properties", Fields: []types.Field{{Name: "BBox", Doc: "bounding box in world coords (Axis-Aligned Bounding Box = AABB)"}, {Name: "VelBBox", Doc: "velocity-projected bounding box in world coords: extend BBox to include future position of moving bodies -- collision must be made on this basis"}, {Name: "BSphere", Doc: "bounding sphere in local coords"}, {Name: "Area", Doc: "area"}, {Name: "Volume", Doc: "volume"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Body", IDName: "body", Doc: "Body is the common interface for all body types"})

// BodyBaseType is the [types.Type] for [BodyBase]
var BodyBaseType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.BodyBase", IDName: "body-base", Doc: "BodyBase is the base type for all specific Body types", Embeds: []types.Field{{Name: "NodeBase"}}, Fields: []types.Field{{Name: "Rigid", Doc: "rigid body properties, including mass, bounce, friction etc"}, {Name: "Vis", Doc: "visualization name -- looks up an entry in the scene library that provides the visual representation of this body"}, {Name: "Color", Doc: "default color of body for basic InitLibrary configuration"}}, Instance: &BodyBase{}})

// NewBodyBase returns a new [BodyBase] with the given optional parent:
// BodyBase is the base type for all specific Body types
func NewBodyBase(parent ...tree.Node) *BodyBase { return tree.New[*BodyBase](parent...) }

// NodeType returns the [*types.Type] of [BodyBase]
func (t *BodyBase) NodeType() *types.Type { return BodyBaseType }

// New returns a new [*BodyBase] value
func (t *BodyBase) New() tree.Node { return &BodyBase{} }

// SetRigid sets the [BodyBase.Rigid]:
// rigid body properties, including mass, bounce, friction etc
func (t *BodyBase) SetRigid(v Rigid) *BodyBase { t.Rigid = v; return t }

// SetVis sets the [BodyBase.Vis]:
// visualization name -- looks up an entry in the scene library that provides the visual representation of this body
func (t *BodyBase) SetVis(v string) *BodyBase { t.Vis = v; return t }

// SetColor sets the [BodyBase.Color]:
// default color of body for basic InitLibrary configuration
func (t *BodyBase) SetColor(v string) *BodyBase { t.Color = v; return t }

// SetInitial sets the [BodyBase.Initial]
func (t *BodyBase) SetInitial(v State) *BodyBase { t.Initial = v; return t }

// SetRel sets the [BodyBase.Rel]
func (t *BodyBase) SetRel(v State) *BodyBase { t.Rel = v; return t }

// BoxType is the [types.Type] for [Box]
var BoxType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Box", IDName: "box", Doc: "Box is a box body shape", Embeds: []types.Field{{Name: "BodyBase"}}, Fields: []types.Field{{Name: "Size", Doc: "size of box in each dimension (units arbitrary, as long as they are all consistent -- meters is typical)"}}, Instance: &Box{}})

// NewBox returns a new [Box] with the given optional parent:
// Box is a box body shape
func NewBox(parent ...tree.Node) *Box { return tree.New[*Box](parent...) }

// NodeType returns the [*types.Type] of [Box]
func (t *Box) NodeType() *types.Type { return BoxType }

// New returns a new [*Box] value
func (t *Box) New() tree.Node { return &Box{} }

// SetSize sets the [Box.Size]:
// size of box in each dimension (units arbitrary, as long as they are all consistent -- meters is typical)
func (t *Box) SetSize(v math32.Vector3) *Box { t.Size = v; return t }

// SetInitial sets the [Box.Initial]
func (t *Box) SetInitial(v State) *Box { t.Initial = v; return t }

// SetRel sets the [Box.Rel]
func (t *Box) SetRel(v State) *Box { t.Rel = v; return t }

// SetRigid sets the [Box.Rigid]
func (t *Box) SetRigid(v Rigid) *Box { t.Rigid = v; return t }

// SetVis sets the [Box.Vis]
func (t *Box) SetVis(v string) *Box { t.Vis = v; return t }

// SetColor sets the [Box.Color]
func (t *Box) SetColor(v string) *Box { t.Color = v; return t }

// CapsuleType is the [types.Type] for [Capsule]
var CapsuleType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Capsule", IDName: "capsule", Doc: "Capsule is a generalized cylinder body shape, with hemispheres at each end,\nwith separate radii for top and bottom.", Embeds: []types.Field{{Name: "BodyBase"}}, Fields: []types.Field{{Name: "Height", Doc: "height of the cylinder portion of the capsule"}, {Name: "TopRad", Doc: "radius of the top hemisphere"}, {Name: "BotRad", Doc: "radius of the bottom hemisphere"}}, Instance: &Capsule{}})

// NewCapsule returns a new [Capsule] with the given optional parent:
// Capsule is a generalized cylinder body shape, with hemispheres at each end,
// with separate radii for top and bottom.
func NewCapsule(parent ...tree.Node) *Capsule { return tree.New[*Capsule](parent...) }

// NodeType returns the [*types.Type] of [Capsule]
func (t *Capsule) NodeType() *types.Type { return CapsuleType }

// New returns a new [*Capsule] value
func (t *Capsule) New() tree.Node { return &Capsule{} }

// SetHeight sets the [Capsule.Height]:
// height of the cylinder portion of the capsule
func (t *Capsule) SetHeight(v float32) *Capsule { t.Height = v; return t }

// SetTopRad sets the [Capsule.TopRad]:
// radius of the top hemisphere
func (t *Capsule) SetTopRad(v float32) *Capsule { t.TopRad = v; return t }

// SetBotRad sets the [Capsule.BotRad]:
// radius of the bottom hemisphere
func (t *Capsule) SetBotRad(v float32) *Capsule { t.BotRad = v; return t }

// SetInitial sets the [Capsule.Initial]
func (t *Capsule) SetInitial(v State) *Capsule { t.Initial = v; return t }

// SetRel sets the [Capsule.Rel]
func (t *Capsule) SetRel(v State) *Capsule { t.Rel = v; return t }

// SetRigid sets the [Capsule.Rigid]
func (t *Capsule) SetRigid(v Rigid) *Capsule { t.Rigid = v; return t }

// SetVis sets the [Capsule.Vis]
func (t *Capsule) SetVis(v string) *Capsule { t.Vis = v; return t }

// SetColor sets the [Capsule.Color]
func (t *Capsule) SetColor(v string) *Capsule { t.Color = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Contact", IDName: "contact", Doc: "Contact is one pairwise point of contact between two bodies.\nContacts are represented in spherical terms relative to the\nspherical BBox of A and B.", Fields: []types.Field{{Name: "A", Doc: "one body"}, {Name: "B", Doc: "the other body"}, {Name: "NormB", Doc: "normal pointing from center of B to center of A"}, {Name: "PtB", Doc: "point on spherical shell of B where A is contacting"}, {Name: "Dist", Doc: "distance from PtB along NormB to contact point on spherical shell of A"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Contacts", IDName: "contacts", Doc: "Contacts is a slice list of contacts"})

// CylinderType is the [types.Type] for [Cylinder]
var CylinderType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Cylinder", IDName: "cylinder", Doc: "Cylinder is a generalized cylinder body shape, with separate radii for top and bottom.\nA cone has a zero radius at one end.", Embeds: []types.Field{{Name: "BodyBase"}}, Fields: []types.Field{{Name: "Height", Doc: "height of the cylinder"}, {Name: "TopRad", Doc: "radius of the top -- set to 0 for a cone"}, {Name: "BotRad", Doc: "radius of the bottom"}}, Instance: &Cylinder{}})

// NewCylinder returns a new [Cylinder] with the given optional parent:
// Cylinder is a generalized cylinder body shape, with separate radii for top and bottom.
// A cone has a zero radius at one end.
func NewCylinder(parent ...tree.Node) *Cylinder { return tree.New[*Cylinder](parent...) }

// NodeType returns the [*types.Type] of [Cylinder]
func (t *Cylinder) NodeType() *types.Type { return CylinderType }

// New returns a new [*Cylinder] value
func (t *Cylinder) New() tree.Node { return &Cylinder{} }

// SetHeight sets the [Cylinder.Height]:
// height of the cylinder
func (t *Cylinder) SetHeight(v float32) *Cylinder { t.Height = v; return t }

// SetTopRad sets the [Cylinder.TopRad]:
// radius of the top -- set to 0 for a cone
func (t *Cylinder) SetTopRad(v float32) *Cylinder { t.TopRad = v; return t }

// SetBotRad sets the [Cylinder.BotRad]:
// radius of the bottom
func (t *Cylinder) SetBotRad(v float32) *Cylinder { t.BotRad = v; return t }

// SetInitial sets the [Cylinder.Initial]
func (t *Cylinder) SetInitial(v State) *Cylinder { t.Initial = v; return t }

// SetRel sets the [Cylinder.Rel]
func (t *Cylinder) SetRel(v State) *Cylinder { t.Rel = v; return t }

// SetRigid sets the [Cylinder.Rigid]
func (t *Cylinder) SetRigid(v Rigid) *Cylinder { t.Rigid = v; return t }

// SetVis sets the [Cylinder.Vis]
func (t *Cylinder) SetVis(v string) *Cylinder { t.Vis = v; return t }

// SetColor sets the [Cylinder.Color]
func (t *Cylinder) SetColor(v string) *Cylinder { t.Color = v; return t }

// GroupType is the [types.Type] for [Group]
var GroupType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Group", IDName: "group", Doc: "Group is a container of bodies, joints, or other groups\nit should be used strategically to partition the space\nand its BBox is used to optimize tree-based collision detection.\nUse a group for the top-level World node as well.", Embeds: []types.Field{{Name: "NodeBase"}}, Instance: &Group{}})

// NewGroup returns a new [Group] with the given optional parent:
// Group is a container of bodies, joints, or other groups
// it should be used strategically to partition the space
// and its BBox is used to optimize tree-based collision detection.
// Use a group for the top-level World node as well.
func NewGroup(parent ...tree.Node) *Group { return tree.New[*Group](parent...) }

// NodeType returns the [*types.Type] of [Group]
func (t *Group) NodeType() *types.Type { return GroupType }

// New returns a new [*Group] value
func (t *Group) New() tree.Node { return &Group{} }

// SetInitial sets the [Group.Initial]
func (t *Group) SetInitial(v State) *Group { t.Initial = v; return t }

// SetRel sets the [Group.Rel]
func (t *Group) SetRel(v State) *Group { t.Rel = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.BodyPoint", IDName: "body-point", Doc: "BodyPoint contains a Body and a Point on that body", Fields: []types.Field{{Name: "Body"}, {Name: "Point"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Node", IDName: "node", Doc: "Node is the common interface for all nodes"})

// NodeBaseType is the [types.Type] for [NodeBase]
var NodeBaseType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.NodeBase", IDName: "node-base", Doc: "NodeBase is the basic node, which has position, rotation, velocity\nand computed bounding boxes, etc.\nThere are only three different kinds of Nodes: Group, Body, and Joint", Embeds: []types.Field{{Name: "NodeBase"}}, Fields: []types.Field{{Name: "Initial", Doc: "initial position, orientation, velocity in *local* coordinates (relative to parent)"}, {Name: "Rel", Doc: "current relative (local) position, orientation, velocity -- only change these values, as abs values are computed therefrom"}, {Name: "Abs", Doc: "current absolute (world) position, orientation, velocity"}, {Name: "BBox", Doc: "bounding box in world coordinates (aggregated for groups)"}}, Instance: &NodeBase{}})

// NewNodeBase returns a new [NodeBase] with the given optional parent:
// NodeBase is the basic node, which has position, rotation, velocity
// and computed bounding boxes, etc.
// There are only three different kinds of Nodes: Group, Body, and Joint
func NewNodeBase(parent ...tree.Node) *NodeBase { return tree.New[*NodeBase](parent...) }

// NodeType returns the [*types.Type] of [NodeBase]
func (t *NodeBase) NodeType() *types.Type { return NodeBaseType }

// New returns a new [*NodeBase] value
func (t *NodeBase) New() tree.Node { return &NodeBase{} }

// SetInitial sets the [NodeBase.Initial]:
// initial position, orientation, velocity in *local* coordinates (relative to parent)
func (t *NodeBase) SetInitial(v State) *NodeBase { t.Initial = v; return t }

// SetRel sets the [NodeBase.Rel]:
// current relative (local) position, orientation, velocity -- only change these values, as abs values are computed therefrom
func (t *NodeBase) SetRel(v State) *NodeBase { t.Rel = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.NodeFlags", IDName: "node-flags", Doc: "NodeFlags define node bitflags -- uses ki Flags field (64 bit capacity)"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Rigid", IDName: "rigid", Doc: "Rigid contains the full specification of a given object's basic physics\nproperties including position, orientation, velocity.  These", Fields: []types.Field{{Name: "InvMass", Doc: "1/mass -- 0 for no mass"}, {Name: "Bounce", Doc: "COR or coefficient of restitution -- how elastic is the collision i.e., final velocity / initial velocity"}, {Name: "Friction", Doc: "friction coefficient -- how much friction is generated by transverse motion"}, {Name: "Force", Doc: "record of computed force vector from last iteration"}, {Name: "RotInertia", Doc: "Last calculated rotational inertia matrix in local coords"}}})

// SphereType is the [types.Type] for [Sphere]
var SphereType = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.Sphere", IDName: "sphere", Doc: "Sphere is a spherical body shape.", Embeds: []types.Field{{Name: "BodyBase"}}, Fields: []types.Field{{Name: "Radius", Doc: "radius"}}, Instance: &Sphere{}})

// NewSphere returns a new [Sphere] with the given optional parent:
// Sphere is a spherical body shape.
func NewSphere(parent ...tree.Node) *Sphere { return tree.New[*Sphere](parent...) }

// NodeType returns the [*types.Type] of [Sphere]
func (t *Sphere) NodeType() *types.Type { return SphereType }

// New returns a new [*Sphere] value
func (t *Sphere) New() tree.Node { return &Sphere{} }

// SetRadius sets the [Sphere.Radius]:
// radius
func (t *Sphere) SetRadius(v float32) *Sphere { t.Radius = v; return t }

// SetInitial sets the [Sphere.Initial]
func (t *Sphere) SetInitial(v State) *Sphere { t.Initial = v; return t }

// SetRel sets the [Sphere.Rel]
func (t *Sphere) SetRel(v State) *Sphere { t.Rel = v; return t }

// SetRigid sets the [Sphere.Rigid]
func (t *Sphere) SetRigid(v Rigid) *Sphere { t.Rigid = v; return t }

// SetVis sets the [Sphere.Vis]
func (t *Sphere) SetVis(v string) *Sphere { t.Vis = v; return t }

// SetColor sets the [Sphere.Color]
func (t *Sphere) SetColor(v string) *Sphere { t.Color = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/xyz/physics.State", IDName: "state", Doc: "State contains the basic physical state including position, orientation, velocity.\nThese are only the values that can be either relative or absolute -- other physical\nstate values such as Mass should go in Rigid.", Fields: []types.Field{{Name: "Pos", Doc: "position of center of mass of object"}, {Name: "Quat", Doc: "rotation specified as a Quat"}, {Name: "LinVel", Doc: "linear velocity"}, {Name: "AngVel", Doc: "angular velocity"}}})
