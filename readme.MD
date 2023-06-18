# Configuração e Execução

Este arquivo contém instruções para configurar e executar o aplicativo, incluindo a inicialização do banco de dados no Docker, a execução dos testes e a execução do arquivo principal (`main.go`).

## Pré-requisitos

Antes de começar, verifique se você tem os seguintes pré-requisitos instalados em sua máquina:

- Docker: [Instruções de instalação do Docker](https://docs.docker.com/get-docker/)
## Configuração e Instalação

1. Clone este repositório em sua máquina local:

```bash
git clone <git@github.com:VictorOliveiraPy/data-manipulation.git>
cd <NOME_DO_DIRETORIO>
```

Para buildar o projeto e executar o aplicativo, siga as etapas abaixo:

Execute o seguinte comando para buildar o projeto:
```bash
make build
````
Em seguida, execute o comando a seguir para iniciar o aplicativo

```bash
make run
````


## Execução dos Testes

Para executar os testes, utilize o seguinte comando:

```bash
make test
```

## Estrutura do Banco de Dados


Para lidar com os dados do arquivo TXT, foram criadas duas tabelas no banco de dados:

1. **Tabela "raw_client_data"**: Esta tabela foi criada para armazenar os dados brutos sem formatação. Ela possui colunas correspondentes aos campos presentes no arquivo TXT, como "document" para armazenar CPF/CNPJ, dados privados, datas de última compra, valor médio do ticket, loja mais frequente, etc. Todos os campos nessa tabela foram definidos como strings para acomodar os dados brutos.

2. **Tabela "client_data"**: Essa tabela foi criada para armazenar os dados após a formatação e correção de tipagem. Ela possui colunas com os tipos de dados apropriados para cada campo, como "document" para armazenar CPF/CNPJ no formato adequado, datas no formato de data/hora, números para valores monetários, etc. Além disso, foi adicionada uma coluna chamada "document_type" na tabela "client_data", que indica o tipo de documento (CPF ou CNPJ) para cada registro armazenado.

Decidi separar os dados brutos dos dados formatados em tabelas distintas para manter a integridade dos dados inseridos e preservar um registro original dos dados, o que pode ser útil para fins de auditoria. A tabela ```'raw_client_data'``` é responsável por armazenar os dados brutos sem formatação, mantendo um histórico fiel do que foi importado. Isso permite rastrear a origem dos dados e ter uma visão clara dos registros iniciais.
Por outro lado, a tabela `````'client_data'````` armazena os dados formatados e corrigidos. Essa separação facilita a visualização e o gerenciamento dos dados, pois concentra as informações prontas para consulta e análise. Além disso, a tabela 'client_data' pode ser otimizada especificamente para consultas.
Essa abordagem também simplifica a manutenção do banco de dados, pois é possível direcionar ações específicas para cada tabela. Por exemplo, é mais fácil executar processos de limpeza e transformação de dados na tabela 'raw_client_data' sem interferir nos dados já formatados e prontos para uso na tabela ```'client_data'```.