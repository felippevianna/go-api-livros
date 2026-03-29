# 🚀 Roadmap de Evolução: API de Livros (Skoob Clone)

Este documento lista as próximas funcionalidades e melhorias técnicas planejadas para elevar o nível de maturidade do projeto.

---

## 🛠 Fase 1: Finalização do Core CRUD
- [x] **Implementar Update (PUT/PATCH):** Permitir a edição de informações de um livro (ex: mudar status de leitura ou corrigir título).
- [x] **Soft Delete:** Alterar a exclusão física para exclusão lógica (preservando dados no banco com um campo `deleted_at`).

## 📊 Fase 2: Relacionamentos e Complexidade de Dados
- [x] **Entidade de Autores:** Criar uma tabela própria para `Autores` e relacioná-la com `Livros` (Um autor tem muitos livros).
- [x] **Categorias/Gêneros:** Implementar um relacionamento de muitos-para-muitos (`Many-to-Many`) para categorias (Ficção, Terror, Biografia).
- [x] **Avaliações e Resenhas:** Permitir que usuários deem notas (1-5 estrelas) e escrevam comentários sobre os livros.

## 🔐 Fase 3: Segurança e Autenticação
- [ ] **Sistema de Usuários:** Cadastro e Login de usuários.
- [ ] **Autenticação JWT:** Proteger as rotas de escrita (Post/Put/Delete) para que apenas usuários autenticados possam alterar dados.
- [ ] **Middleware de Permissão:** Garantir que um usuário só possa editar/deletar os livros da sua própria estante.

## 🌐 Fase 4: Integrações e Experiência do Usuário
- [ ] **Integração Google Books API:** Criar um endpoint que busca dados automáticos de um livro pelo ISBN ou Título.
- [ ] **Upload de Imagens:** Permitir o upload e armazenamento de capas de livros (usando S3 ou local storage).
- [ ] **Busca Avançada:** Implementar filtros por autor, título e ano de publicação diretamente na URL.

## 📖 Fase 5: Qualidade e Documentação
- [ ] **Swagger UI:** Gerar documentação interativa da API para facilitar o consumo por aplicações Front-end.
- [ ] **Testes Unitários:** Implementar testes para a camada de Service e Repository.
- [ ] **Logs e Monitoramento:** Adicionar logs estruturados para rastrear erros em produção.

## ⚙️ Fase 6: Possíveis melhorias
- [ ] **Mecanismo de Retry no DB:** Implementar um loop de reconexão no `config/db.go` para evitar que a API caia caso o banco de dados demore a iniciar no Docker.
- [ ] **Padrão DTO (Data Transfer Objects):** Separar as structs de entrada da API (Request) das structs do Banco de Dados (Model) para ter mais controle sobre o que o usuário envia.
- [ ] **Cálculo de Médias (Aggregates):** Criar um endpoint ou campo calculado que retorne a nota média de um livro baseada nas avaliações.
- [ ] **Tratamento Global de Erros:** Criar um middleware para padronizar as respostas de erro da API, evitando repetição de código nos Handlers.
- [ ] **Validações Customizadas:** Utilizar a biblioteca `validator/v10` para garantir regras de negócio direto nas structs (ex: impedir notas maiores que 5 ou títulos vazios).