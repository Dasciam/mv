package packet

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	RecipeShapeless int32 = iota
	RecipeShaped
	RecipeFurnace
	RecipeFurnaceData
	RecipeMulti
	RecipeShulkerBox
	RecipeShapelessChemistry
	RecipeShapedChemistry
	RecipeSmithingTransform
	RecipeSmithingTrim
)

// lookupRecipeType looks up the recipe type for a Recipe. False is returned if
// none was found.
func lookupRecipeType(x protocol.Recipe, recipeType *int32) bool {
	switch x.(type) {
	case *ShapelessRecipe:
		*recipeType = RecipeShapeless
	case *ShapedRecipe:
		*recipeType = RecipeShaped
	case *FurnaceRecipe:
		*recipeType = RecipeFurnace
	case *FurnaceDataRecipe:
		*recipeType = RecipeFurnaceData
	case *MultiRecipe:
		*recipeType = RecipeMulti
	case *ShulkerBoxRecipe:
		*recipeType = RecipeShulkerBox
	case *ShapelessChemistryRecipe:
		*recipeType = RecipeShapelessChemistry
	case *ShapedChemistryRecipe:
		*recipeType = RecipeShapedChemistry
	case *SmithingTransformRecipe:
		*recipeType = RecipeSmithingTransform
	case *SmithingTrimRecipe:
		*recipeType = RecipeSmithingTrim
	default:
		return false
	}
	return true
}

// ShapelessRecipe is a recipe that has no particular shape. Its functionality is shared with the
// RecipeShulkerBox and RecipeShapelessChemistry types.
type ShapelessRecipe struct {
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	RecipeID string
	// Input is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output.
	Input []protocol.ItemDescriptorCount
	// Output is a list of items that are created as a result of crafting the recipe.
	Output []protocol.ItemStack
	// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can always be empty.
	UUID uuid.UUID
	// Block is the block name that is required to craft the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'crafting_table' as an example.
	// The available blocks are:
	// - crafting_table
	// - cartography_table
	// - stonecutter
	// - furnace
	// - blast_furnace
	// - smoker
	// - campfire
	Block string
	// Priority ...
	Priority int32
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
}

// ShulkerBoxRecipe is a shapeless recipe made specifically for shulker box crafting, so that they don't lose
// their user data when dyeing a shulker box.
type ShulkerBoxRecipe struct {
	ShapelessRecipe
}

// ShapelessChemistryRecipe is a recipe specifically made for chemistry related features, which exist only in
// the Education Edition. They function the same as shapeless recipes do.
type ShapelessChemistryRecipe struct {
	ShapelessRecipe
}

// ShapedRecipe is a recipe that has a specific shape that must be used to craft the output of the recipe.
// Trying to craft the item in any other shape will not work. The ShapedRecipe is of the same structure as the
// ShapedChemistryRecipe.
type ShapedRecipe struct {
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	RecipeID string
	// Width is the width of the recipe's shape.
	Width int32
	// Height is the height of the recipe's shape.
	Height int32
	// Input is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output. The amount of input items must be exactly equal to Width * Height.
	Input []protocol.ItemDescriptorCount
	// Output is a list of items that are created as a result of crafting the recipe.
	Output []protocol.ItemStack
	// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can always be empty.
	UUID uuid.UUID
	// Block is the block name that is required to craft the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'crafting_table' as an example.
	Block string
	// Priority ...
	Priority int32
	// AssumeSymmetry specifies if the recipe is symmetrical. If this is set to true, the recipe will be
	// mirrored along the diagonal axis. This means that the recipe will be the same if rotated 180 degrees.
	AssumeSymmetry bool
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
}

// ShapedChemistryRecipe is a recipe specifically made for chemistry related features, which exist only in the
// Education Edition. It functions the same as a normal ShapedRecipe.
type ShapedChemistryRecipe struct {
	ShapedRecipe
}

// FurnaceRecipe is a recipe that is specifically used for all kinds of furnaces. These recipes don't just
// apply to furnaces, but also blast furnaces and smokers.
type FurnaceRecipe struct {
	// InputType is the item type of the input item. The metadata value of the item is not used in the
	// FurnaceRecipe. Use FurnaceDataRecipe to allow an item with only one metadata value.
	InputType protocol.ItemType
	// Output is the item that is created as a result of smelting/cooking an item in the furnace.
	Output protocol.ItemStack
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'furnace' as an example.
	Block string
}

// FurnaceDataRecipe is a recipe specifically used for furnace-type crafting stations. It is equal to
// FurnaceRecipe, except it has an input item with a specific metadata value, instead of any metadata value.
type FurnaceDataRecipe struct {
	FurnaceRecipe
}

// MultiRecipe serves as an 'enable' switch for multi-shape recipes.
type MultiRecipe struct {
	// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can always be empty.
	UUID uuid.UUID
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
}

// SmithingTransformRecipe is a recipe specifically used for smithing tables. It has three input items and adds them
// together, resulting in a new item.
type SmithingTransformRecipe struct {
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	RecipeID string
	// Template is the item that is used to shape the Base item based on the Addition being applied.
	Template protocol.ItemDescriptorCount
	// Base is the item that the Addition is being applied to in the smithing table.
	Base protocol.ItemDescriptorCount
	// Addition is the item that is being added to the Base item to result in a modified item.
	Addition protocol.ItemDescriptorCount
	// Result is the resulting item from the two items being added together.
	Result protocol.ItemStack
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed with
	// 'minecraft:', so it will look like 'smithing_table' as an example.
	Block string
}

// SmithingTrimRecipe is a recipe specifically used for applying armour trims to an armour piece inside a smithing table.
type SmithingTrimRecipe struct {
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	RecipeID string
	// Template is the item that is used to shape the Base item based on the Addition being applied.
	Template protocol.ItemDescriptorCount
	// Base is the item that the Addition is being applied to in the smithing table.
	Base protocol.ItemDescriptorCount
	// Addition is the item that is being added to the Base item to result in a modified item.
	Addition protocol.ItemDescriptorCount
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed with
	// 'minecraft:', so it will look like 'smithing_table' as an example.
	Block string
}

// Marshal ...
func (recipe *ShapelessRecipe) Marshal(w *protocol.Writer) {
	marshalShapeless(w, recipe)
}

// Unmarshal ...
func (recipe *ShapelessRecipe) Unmarshal(r *protocol.Reader) {
	marshalShapeless(r, recipe)
}

// Marshal ...
func (recipe *ShulkerBoxRecipe) Marshal(w *protocol.Writer) {
	marshalShapeless(w, &recipe.ShapelessRecipe)
}

// Unmarshal ...
func (recipe *ShulkerBoxRecipe) Unmarshal(r *protocol.Reader) {
	marshalShapeless(r, &recipe.ShapelessRecipe)
}

// Marshal ...
func (recipe *ShapelessChemistryRecipe) Marshal(w *protocol.Writer) {
	marshalShapeless(w, &recipe.ShapelessRecipe)
}

// Unmarshal ...
func (recipe *ShapelessChemistryRecipe) Unmarshal(r *protocol.Reader) {
	marshalShapeless(r, &recipe.ShapelessRecipe)
}

// Marshal ...
func (recipe *ShapedRecipe) Marshal(w *protocol.Writer) {
	marshalShaped(w, recipe)
}

// Unmarshal ...
func (recipe *ShapedRecipe) Unmarshal(r *protocol.Reader) {
	marshalShaped(r, recipe)
}

// Marshal ...
func (recipe *ShapedChemistryRecipe) Marshal(w *protocol.Writer) {
	marshalShaped(w, &recipe.ShapedRecipe)
}

// Unmarshal ...
func (recipe *ShapedChemistryRecipe) Unmarshal(r *protocol.Reader) {
	marshalShaped(r, &recipe.ShapedRecipe)
}

// Marshal ...
func (recipe *FurnaceRecipe) Marshal(w *protocol.Writer) {
	w.Varint32(&recipe.InputType.NetworkID)
	w.Item(&recipe.Output)
	w.String(&recipe.Block)
}

// Unmarshal ...
func (recipe *FurnaceRecipe) Unmarshal(r *protocol.Reader) {
	r.Varint32(&recipe.InputType.NetworkID)
	r.Item(&recipe.Output)
	r.String(&recipe.Block)
}

// Marshal ...
func (recipe *FurnaceDataRecipe) Marshal(w *protocol.Writer) {
	w.Varint32(&recipe.InputType.NetworkID)
	aux := int32(recipe.InputType.MetadataValue)
	w.Varint32(&aux)
	w.Item(&recipe.Output)
	w.String(&recipe.Block)
}

// Unmarshal ...
func (recipe *FurnaceDataRecipe) Unmarshal(r *protocol.Reader) {
	var dataValue int32
	r.Varint32(&recipe.InputType.NetworkID)
	r.Varint32(&dataValue)
	recipe.InputType.MetadataValue = uint32(dataValue)
	r.Item(&recipe.Output)
	r.String(&recipe.Block)
}

// Marshal ...
func (recipe *MultiRecipe) Marshal(w *protocol.Writer) {
	w.UUID(&recipe.UUID)
	w.Varuint32(&recipe.RecipeNetworkID)
}

// Unmarshal ...
func (recipe *MultiRecipe) Unmarshal(r *protocol.Reader) {
	r.UUID(&recipe.UUID)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// Marshal ...
func (recipe *SmithingTransformRecipe) Marshal(w *protocol.Writer) {
	w.String(&recipe.RecipeID)
	w.ItemDescriptorCount(&recipe.Template)
	w.ItemDescriptorCount(&recipe.Base)
	w.ItemDescriptorCount(&recipe.Addition)
	w.Item(&recipe.Result)
	w.String(&recipe.Block)
	w.Varuint32(&recipe.RecipeNetworkID)
}

// Unmarshal ...
func (recipe *SmithingTransformRecipe) Unmarshal(r *protocol.Reader) {
	r.String(&recipe.RecipeID)
	r.ItemDescriptorCount(&recipe.Template)
	r.ItemDescriptorCount(&recipe.Base)
	r.ItemDescriptorCount(&recipe.Addition)
	r.Item(&recipe.Result)
	r.String(&recipe.Block)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// Marshal ...
func (recipe *SmithingTrimRecipe) Marshal(w *protocol.Writer) {
	w.String(&recipe.RecipeID)
	w.ItemDescriptorCount(&recipe.Template)
	w.ItemDescriptorCount(&recipe.Base)
	w.ItemDescriptorCount(&recipe.Addition)
	w.String(&recipe.Block)
	w.Varuint32(&recipe.RecipeNetworkID)
}

// Unmarshal ...
func (recipe *SmithingTrimRecipe) Unmarshal(r *protocol.Reader) {
	r.String(&recipe.RecipeID)
	r.ItemDescriptorCount(&recipe.Template)
	r.ItemDescriptorCount(&recipe.Base)
	r.ItemDescriptorCount(&recipe.Addition)
	r.String(&recipe.Block)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShaped ...
func marshalShaped(r protocol.IO, recipe *ShapedRecipe) {
	r.String(&recipe.RecipeID)
	r.Varint32(&recipe.Width)
	r.Varint32(&recipe.Height)
	protocol.FuncSliceOfLen(r, uint32(recipe.Width*recipe.Height), &recipe.Input, r.ItemDescriptorCount)
	protocol.FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	r.Bool(&recipe.AssumeSymmetry)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShapeless ...
func marshalShapeless(r protocol.IO, recipe *ShapelessRecipe) {
	r.String(&recipe.RecipeID)
	protocol.FuncSlice(r, &recipe.Input, r.ItemDescriptorCount)
	protocol.FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	r.Varuint32(&recipe.RecipeNetworkID)
}