package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Структура, которую мы будем создавать с помощью шаблона Строитель
type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

// Интерфейс для создания пользователя
type UserBuilderI interface {
	SetFirstName(firstName string) UserBuilderI
	SetLastName(lastName string) UserBuilderI
	SetEmail(email string) UserBuilderI
	SetAge(age int) UserBuilderI
	Build() User
}

// Реализация интерфейса UserBuilder
type UserBuilder struct {
	user User
}

func NewUserBuilder() UserBuilderI {
	return &UserBuilder{}
}

func (b *UserBuilder) SetFirstName(firstName string) UserBuilderI {
	b.user.FirstName = firstName
	return b
}

func (b *UserBuilder) SetLastName(lastName string) UserBuilderI {
	b.user.LastName = lastName
	return b
}

func (b *UserBuilder) SetEmail(email string) UserBuilderI {
	b.user.Email = email
	return b
}

func (b *UserBuilder) SetAge(age int) UserBuilderI {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

// Интерфейс для Директора
type UserDirector interface {
	CreateUser(firstName, lastName, email string, age int) User
}

// Реализация интерфейса UserDirector
type userDirector struct {
	builder UserBuilderI
}

func NewUserDirector(builder UserBuilderI) UserDirector {
	return &userDirector{builder: builder}
}

func (d *userDirector) CreateUser(firstName, lastName, email string, age int) User {
	return d.builder.
		SetFirstName(firstName).
		SetLastName(lastName).
		SetEmail(email).
		SetAge(age).
		Build()
}

/*Пример использования шаблона Строитель
func main() {
	userBuilder := NewUserBuilder()
	user := userBuilder.SetFirstName("Name").
		SetLastName("Surname").
		SetEmail("ilia@mail.ru").
		SetAge(21).
		Build()

	fmt.Printf("%+v\n", user)

	userDirector := NewUserDirector(userBuilder)
	user = userDirector.CreateUser("Name2", "Surname2", "Boris@gmail.com", 30)

	fmt.Printf("%+v\n", user)
}
Применимость паттерна "Строитель":
Использование паттерна рекомендуется, когда процесс инстанцирования объекта состоит из множества шагов, и эти шаги должны выполняться в определенной последовательности.
Паттерн "Строитель" целесообразно применять, когда имеются объекты с большим количеством опциональных параметров, и чтобы избежать создания непомерно большого количества конструкторов.
Паттерн часто используют для инкапсуляции и повторного использования процесса создания сложного объекта.
Плюсы паттерна "Строитель":
Позволяет изменять внутреннее представление продукта без изменения клиентского кода.
Изолирует сложный код создания продукта от его основной бизнес-логики.
Предоставляет более гибкую и масштабируемую архитектуру, так как позволяет производить различные представления объекта использование того же строительного процесса.
Минусы паттерна "Строитель":
Усложнение кода за счет необходимости создавать дополнительные классы "Строителей".
В некоторых случаях может привести к избыточности или напрасной сложности, особенно если объект простой и не требует последовательного конструирования.
*/
