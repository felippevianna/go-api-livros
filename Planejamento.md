# 🚀 Roadmap de Evolução: API de Livros (Skoob Clone)

Este documento lista as próximas funcionalidades e melhorias técnicas planejadas para elevar o nível de maturidade do projeto.

---

## 🛠 Fase 1: Finalização do Core CRUD
- [ ] **Implementar Update (PUT/PATCH):** Permitir a edição de informações de um livro (ex: mudar status de leitura ou corrigir título).
- [ ] **Soft Delete:** Alterar a exclusão física para exclusão lógica (preservando dados no banco com um campo `deleted_at`).

## 📊 Fase 2: Relacionamentos e Complexidade de Dados
- [ ] **Entidade de Autores:** Criar uma tabela própria para `Autores` e relacioná-la com `Livros` (Um autor tem muitos livros).
- [ ] **Categorias/Gêneros:** Implementar um relacionamento de muitos-para-muitos (`Many-to-Many`) para categorias (Ficção, Terror, Biografia).
- [ ] **Avaliações e Resenhas:** Permitir que usuários deem notas (1-5 estrelas) e escrevam comentários sobre os livros.

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