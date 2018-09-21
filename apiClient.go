package elemeOpenApi

type APIClient struct {
	Message    Message
	Order      Order
	Shop       Shop
	Product    Product
	User       User
	File       File
	Packs      Packs
	Finance    Finance
	Activity   Activity
	Ugc        Ugc
	Market     Market
	Flow       Flow
	Content    Content
	Decoration Decoration
	Card       Card
	Cpc        Cpc
	Sms        Sms
	config     Config
}

type User struct {
	config *Config
}

type Shop struct {
	Setup  Setup
	config *Config
}

type Product struct {
	Item     Item
	Category Category
	Package  Package
	config   *Config
}

type Order struct {
	Delivery Delivery
	config   *Config
}

type Message struct {
	config *Config
}

type File struct {
	config *Config
}

type Item struct {
	config *Config
}

type Category struct {
	config *Config
}

type Packs struct {
	config *Config
}

type Finance struct {
	config *Config
}

type Flash struct {
	config *Config
}

type Shippingfee struct {
	config *Config
}

type Food struct {
	config *Config
}

type Coupon struct {
	config *Config
}

type Activity struct {
	Food        Food
	Flash       Flash
	Shippingfee Shippingfee
	Coupon      Coupon
	config      *Config
}

type Ugc struct {
	config *Config
}

type Market struct {
	config *Config
}

type Flow struct {
	config *Config
}

type Content struct {
	config *Config
}

type Windows struct {
	config *Config
}

type Sign struct {
	config *Config
}

type Poster struct {
	config *Config
}

type Story struct {
	config *Config
}

type Image struct {
	config *Config
}

type Accuratecategory struct {
	config *Config
}

type Burstwindow struct {
	config *Config
}

type Decoration struct {
	Windows          Windows
	Sign             Sign
	Poster           Poster
	Story            Story
	Image            Image
	Accuratecategory Accuratecategory
	Burstwindow      Burstwindow
	config           *Config
}

type Delivery struct {
	config *Config
}

type Setup struct {
	config *Config
}

type Package struct {
	config *Config
}

type Card struct {
	config *Config
}

type Cpc struct {
	config *Config
}

type Sms struct {
	config *Config
}

func NewAPIClient(config Config) APIClient {
	client := APIClient{}
	client.SetConfig(config)
	return client
}

func (client *APIClient) SetConfig(config Config) {
	client.config = config
	client.Message.config = &client.config
	client.Order.config = &client.config
	client.Order.Delivery.config = &client.config
	client.Product.config = &client.config
	client.Product.Item.config = &client.config
	client.Product.Category.config = &client.config
	client.Product.Package.config = &client.config
	client.User.config = &client.config
	client.Shop.config = &client.config
	client.Shop.Setup.config = &client.config
	client.File.config = &client.config
	client.Packs.config = &client.config
	client.Finance.config = &client.config
	client.Activity.config = &client.config
	client.Activity.Flash.config = &client.config
	client.Activity.Shippingfee.config = &client.config
	client.Activity.Food.config = &client.config
	client.Activity.Coupon.config = &client.config
	client.Ugc.config = &client.config
	client.Market.config = &client.config
	client.Flow.config = &client.config
	client.Content.config = &client.config
	client.Decoration.config = &client.config
	client.Decoration.Windows.config = &client.config
	client.Decoration.Sign.config = &client.config
	client.Decoration.Poster.config = &client.config
	client.Decoration.Story.config = &client.config
	client.Decoration.Image.config = &client.config
	client.Decoration.Accuratecategory.config = &client.config
	client.Decoration.Burstwindow.config = &client.config
	client.Card.config = &client.config
	client.Cpc.config = &client.config
	client.Sms.config = &client.config
}
