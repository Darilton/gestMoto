## Resumo

Web app mobile-first  simples e eficiente para uma empresa com 3 sócios gerenciarem motorizadas e motoqueiros, rastreando manutenção, status e finanças como receitas semanais, despesas e saldo atual, com interface intuitiva e segura para uso diário.
## 1. Introdução

Este documento descreve o design e as especificações para um web app mobile-first destinado à gestão financeira e de motorizadas de uma pequena empresa com 3 sócios. A empresa opera com 3 motorizadas e 3 motoqueiros, onde cada motoqueiro gera uma receita semanal específica através de serviços de transporte. O foco é na simplicidade, usabilidade em dispositivos móveis e funcionalidades essenciais para rastrear ativos (motorizadas), histórico de manutenção e aspectos financeiros como receitas, despesas e saldo atual.

O app será desenvolvido como um Progressive Web App (PWA) para garantir compatibilidade com navegadores móveis e desktops, com design responsivo priorizando telas pequenas (mobile-first). Tecnologias sugeridas incluem HTML5, CSS3 (com frameworks como Tailwind CSS para responsividade), JavaScript (com React.js para a interface frontend) e um backend em Node.js com banco de dados PostgreSQL ou MongoDB para armazenamento de dados. Autenticação será implementada via JWT para segurança.

**Objetivos Principais:**

- Facilitar a gestão de motorizadas, incluindo manutenção e atribuição de motoqueiros.
- Permitir visualização e rastreamento de finanças em tempo real.
- Garantir acesso restrito aos 3 sócios, com possibilidade de expansão para motoqueiros visualizarem dados limitados.
- Interface intuitiva, com navegação por menus laterais ou bottom navigation bars otimizados para mobile.

## 2. Requisitos Funcionais

Os requisitos são baseados nas necessidades descritas: gestão de motorizadas e finanças.

### 2.1 Gestão de Motorizadas

- **Cadastro e Edição de Motorizadas:** Cada motorizada deve ter campos como ID único, modelo, placa, ano de fabricação, data de aquisição, status (ativo/inativo) e motoqueiro responsável (selecionável de uma lista de motoqueiros cadastrados).
- **Histórico de Manutenção:** Registro de reparações e trocas de óleo, incluindo data, descrição, custo, fornecedor e quilometragem na ocasião. Deve permitir adicionar novos registros e visualizar histórico em ordem cronológica.
- **Atribuição de Motoqueiros:** Associação de um motoqueiro a uma motorizada. Cada motoqueiro tem perfil com nome, contato, data de contratação e status (ativo/inativo).
- **Visualização de Status:** Dashboard para ver o estado atual de todas as motorizadas, com filtros por status ou motoqueiro.

### 2.2 Gestão Financeira

- **Registro de Receitas:** Entrada semanal de valores gerados por cada motoqueiro/motorizada. Campos incluem data/semana, motoqueiro, valor gerado e observações. Cálculo automático de totais mensais/anuais.
- **Registro de Despesas:** Entrada de gastos, categorizados (ex: combustível, manutenção, salários, impostos). Campos: data, categoria, valor, descrição e comprovante (opção para upload de arquivos).
- **Visualização de Finanças:**
    - Saldo atual: Cálculo em tempo real (receitas totais - despesas totais).
    - Quanto foi feito: Relatórios de receitas por período (semana, mês, ano), por motoqueiro ou total.
    - Quanto é gasto: Relatórios de despesas por categoria e período.
- **Relatórios e Gráficos:** Gráficos simples (usando bibliotecas como Chart.js) para visualizar tendências, como receitas vs. despesas ao longo do tempo.

### 2.3 Funcionalidades Gerais

- **Autenticação e Autorização:** Login com email/senha para os 3 sócios. Perfis de sócios com permissões totais (CRUD em todos os dados). Opção futura para contas de motoqueiros com visualização read-only de suas motorizadas e receitas.
- **Notificações:** Alertas para manutenções pendentes (ex: troca de óleo a cada X km) ou saldos baixos.
- **Exportação de Dados:** Geração de relatórios em PDF ou CSV para finanças e histórico de motorizadas.
- **Busca e Filtros:** Pesquisa por datas, motoqueiros ou categorias em todas as seções.

## 3. Requisitos Não Funcionais

- **Desempenho:** Carregamento rápido em conexões móveis lentas (otimização de imagens e caching via Service Workers para PWA).
- **Segurança:** Criptografia de dados sensíveis (senhas hashed com bcrypt), proteção contra SQL injection e XSS. Backup automático de dados.
- **Usabilidade:** Design mobile-first com telas touch-friendly, ícones intuitivos e suporte a idiomas (português como padrão, com opção para inglês).
- **Escalabilidade:** Projetado para 3 motorizadas/motoqueiros iniciais, mas escalável para mais ativos.
- **Acessibilidade:** Conformidade com WCAG 2.1, incluindo contraste de cores e suporte a leitores de tela.
- **Plataformas Suportadas:** Browsers modernos (Chrome, Safari, Firefox) em Android/iOS e desktops.

## 4. Arquitetura do Sistema

### 4.1 Frontend

- Framework: React.js com React Router para navegação.
- Componentes Principais:
    - Dashboard: Visão geral com cards para motorizadas ativas, saldo atual e receitas semanais.
    - Página de Motorizadas: Lista com detalhes expansíveis, formulários para adicionar/editar.
    - Página de Finanças: Formulários para entradas, gráficos e relatórios.
    - Menu Lateral: Acesso rápido a seções, logout e configurações.
- Estilo: Mobile-first com media queries; temas claros/escuros para melhor usabilidade.

### 4.2 Backend

- API RESTful com endpoints como:
    - /api/motorizadas: GET/POST/PUT/DELETE para gestão de motorizadas.
    - /api/manutencao: POST para adicionar histórico, GET para listar.
    - /api/financas/receitas: POST/GET para receitas.
    - /api/financas/despesas: POST/GET para despesas.
    - /api/relatorios: GET para saldos e gráficos.
- Banco de Dados: Schema relacional:
    - Tabela Motorizadas: id, modelo, placa, ano, status, motoqueiro_id.
    - Tabela Motoqueiros: id, nome, contato, status.
    - Tabela Manutencao: id, motorizada_id, data, tipo (reparacao/oleo), custo, descricao.
    - Tabela Receitas: id, motoqueiro_id, data, valor.
    - Tabela Despesas: id, data, categoria, valor, descricao.
    - Tabela Usuarios: id, email, senha (hash), role (socio).

### 4.3 Integrações

- Armazenamento de Arquivos: Uso de serviços como AWS S3 ou local para uploads de comprovativos.
- Notificações: Integração com Push Notifications via Firebase para alertas mobile.

## 5. Fluxo de Usuário

- **Login:** Usuário (sócio) faz login e é direcionado ao dashboard.
- **Gerenciar Motorizada:** No menu, seleciona "Motorizadas", visualiza lista, clica para editar status ou adicionar manutenção.
- **Registrar Finanças:** Em "Finanças", adiciona receita semanal por motoqueiro ou despesa, visualiza gráficos.
- **Relatórios:** Seleciona período e gera visualizações ou exports.

## 6. Protótipo e Testes

- **Protótipo:** Usar Figma para wireframes mobile-first, com telas como login, dashboard e formulários.
- **Testes:** Unitários (Jest para frontend), integração (Postman para API), e usabilidade com usuários reais (foco em mobile).
- **Implantação:** Hospedagem em Vercel ou Heroku, com domínio personalizado.

## 7. Cronograma Estimado

- Fase 1: Planejamento e Design (2 semanas).
- Fase 2: Desenvolvimento Frontend/Backend (4-6 semanas).
- Fase 3: Testes e Deploy (2 semanas).
- Manutenção: Atualizações mensais baseadas em feedback.
