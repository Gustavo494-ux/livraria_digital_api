package models

// Struct comentadas para serem implementadas corretamente depois.

// type Autor struct {
// 	gorm.Model
// 	Nome           string    `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
// 	DataNascimento time.Time `json:"data_nascimento" gorm:"not null" validate:"required"`
// 	Nacionalidade  string    `json:"nacionalidade" gorm:"not null" validate:"required,min=2,max=100"`
// 	Biografia      string    `json:"biografia" gorm:"type:text"`
// }

// type Editora struct {
// 	gorm.Model
// 	Nome     string `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
// 	Endereco string `json:"endereco" gorm:"not null" validate:"required,min=2,max=255"`
// 	Telefone string `json:"telefone" gorm:"not null" validate:"required,min=10,max=15"`
// 	Email    string `json:"email" gorm:"not null" validate:"required,email"`
// }

// type Categoria struct {
// 	gorm.Model
// 	Nome string `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
// }

// type Livro struct {
// 	gorm.Model
// 	Titulo         string    `json:"titulo" gorm:"not null" validate:"required,min=2,max=100"`
// 	AutorID        uint      `json:"autor_id" gorm:"not null" validate:"required"`     // Chave estrangeira para Autor
// 	EditoraID      uint      `json:"editora_id" gorm:"not null" validate:"required"`   // Chave estrangeira para Editora
// 	CategoriaID    uint      `json:"categoria_id" gorm:"not null" validate:"required"` // Chave estrangeira para Categoria
// 	ISBN           string    `json:"isbn" gorm:"unique;not null" validate:"required,min=10,max=13"`
// 	Preco          float64   `json:"preco" gorm:"not null" validate:"required,min=0"`
// 	DataPublicacao time.Time `json:"data_publicacao" gorm:"not null" validate:"required"`

// 	// Relacionamentos
// 	Autor     Autor     `gorm:"foreignKey:AutorID;references:ID"`     // Relacionamento com Autor
// 	Editora   Editora   `gorm:"foreignKey:EditoraID;references:ID"`   // Relacionamento com Editora
// 	Categoria Categoria `gorm:"foreignKey:CategoriaID;references:ID"` // Relacionamento com Categoria
// }

// type Cargo struct {
// 	gorm.Model
// 	Nome        string  `json:"nome" gorm:"not null" validate:"required,min=2,max=100"` // Nome do cargo (ex: "Gerente", "Vendedor")
// 	Descricao   string  `json:"descricao" gorm:"type:text"`                             // Descrição do cargo (opcional)
// 	SalarioBase float64 `json:"salario_base" gorm:"not null" validate:"required,min=0"` // Salário base para o cargo
// }

// type Funcionario struct {
// 	gorm.Model
// 	Nome            string    `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
// 	CPF             string    `json:"cpf" gorm:"unique;not null" validate:"required,min=11,max=14"`
// 	CargoID         uint      `json:"cargo_id" gorm:"not null" validate:"required"` // Chave estrangeira para Cargo
// 	Cargo           Cargo     `gorm:"foreignKey:CargoID;references:ID"`             // Relacionamento com Cargo
// 	DataContratacao time.Time `json:"data_contratacao" gorm:"not null" validate:"required"`
// 	Telefone        string    `json:"telefone" gorm:"not null" validate:"required,min=10,max=15"`
// 	Email           string    `json:"email" gorm:"not null" validate:"required,email"`
// }
