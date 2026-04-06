// POR QUE: O compilador do Go exige que um "package main"(no arquivo principal) para saber que deve gerar um arquivo executável (.exe ou binário) a partir daqui, e não apenas uma bilbioteca de apoio
// PSEUDO-CÓDIGO: Declaramos que este arquivo é o coração do programa (o pacote principal sempre é o "Package Main")
package main

//POR QUE: Não precisamos reinventar a roda. Trazemos pacotes que nos auxiliam a escrever e executar nosso código(nesse caso de errors e formatação de texto "fmt")
//PSEUDO-CÓDIGO: Importamos ferramentas externas prontas para usarmos no nosso código
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ___PASSO 1. Classe (O molde)___

// POR QUE: Este é o "Molde de papel" no nosso ateliê. Ele define como o objeto será construído na memória RAM quando for instanciado(definido,criado)
// PSEUDO-CÓDIGO: Criamos um novo tipo de dado estruturado (struct) chamado de ProvadorVirtual
type ProvadorVirtual struct {

	//POR QUE: A letra minúscula "p" em pecasVestidas é uma regra do Go para "ENCAPSULAMENTO". Ela proíbe que outros arquivos fora deste pacote toquem nesse Slice diretamente, obgirando o uso dos métodos seguros.
	//PSEUDO-CÓDIGO: Criamos um ATRIBUTO privado chamado "pecasVestidas", que é uma Lista (Slice) restrita para textos(strings)
	pecasVestidas []string
}

// ___PASSO 2. Os Métodos (O comportamento)__

// POR QUE: O "(p *ProvadorVirtual)" é o Method Receiver(pertence a um tipo específico). O "*"(ponteiro) diz à CPU: "Não crie uma cópia do manequim. Vá até o endereço de memória do manequim real "p" e altere ele"
// PSEUDO-CÓDIGO: Criamos o método AdicionarPeca que recebe um texto (peca). Ele pertence exclusivamente ao ProvadorVirtual.
func (p *ProvadorVirtual) AdicionarPeca(peca string) error {

	//POR QUE: Precisamos saber quantas peças já estão no corpo do cliente para aplicar a nova trava de segurança limitadora.
	//PSEUDO-CÓDIGO: Medimos o tamanho(len) do slice "pecasVestidas" para saber quantas peças já estão vestidas no momento.
	tamanho := len(p.pecasVestidas)

	//POR QUE: Esta é um TRAVA DE SEGURANÇA. Se deixarmos empilhar roupas sem limite, a memória do servidor acaba(Panic) e o sistema quebra
	//PSEUDO-CÓDIGO: Verificamos e o tamanho já bateu o limite máximo de 5 peças.
	if tamanho >= 5 {

		//POR QUE: Interrompe a func na hora. Devolvendo o motivo exato do bloqueio
		//PSEUDO-CÓDIGO: Retorna um erro formal avisando que o limite foi atingido.
		return errors.New("Limite de peças atingido! O provador suporta no máximo 5 peças de roupa.")
	}

	//POR QUE: A função "append" é ótimizada pelo Go para alocar mais espaço na memória apenas se a nossa "caixa"(capacidade do slice) estiver cheia. Esta é a operação PUSH(empilhar) da nossa Pilha, com custo O(1) | O(1):
	//PSEUDO-CÓDIGO: Pegamos o slice atual, anexamos a uma nova "peca no final dele, e guardamos o resultado de volta no mesmo lugar.
	p.pecasVestidas = append(p.pecasVestidas, peca)
	fmt.Printf("Sucesso! Você vestiu: %s\n", peca)

	//POR QUE: Como a func exige um retorno tipo erro, e a ação deu certo, retornamos 'nil' (nulo/vazio) para sinalisar sucesso, ou seja, que não houve erro.
	//PSEUDO-CÓDIGO: Retorn 'nil' (sem erros).
	return nil
}

// POR QUE: O retorno multiplo "(string, error)" é o padrão ouro da segurança no Go Se algo der errado, ele avisa, se der certo ele entrega os dados.
// PSEUDO-CÓDIGO: Criamos o método DesfazerLook. Ele não recebe parâmetros, mas devolve Duas coisas: um texto( a peça) e um erro.
func (p *ProvadorVirtual) DesfazerLook() (string, error) {

	//POR QUE: Precisamos saber a quantidade exata para descobrir matematicamente quem
	//PSEUDO-CÓDIGO: Medimos(len) quantos itens existem dentro do slice "pecasVestidas"
	tamanho := len(p.pecasVestidas)

	//POR QUE: Isso é ""Programação Defensiva". Se tentarmo tirar uma peça de um slice vazio, o Go dá um erro fatal de memória (Panic) e derruba o servidor. Para evitar isso, verificamos antes se o slice tem algo dentro.
	//PSEUDO-CÓDIGO: Verificamos se o tamanho é exatamento igual a zero(vazio)
	if tamanho == 0 {

		//POR QUE: O erro "errors.New" cria o objeto de erro que o back-end vai enviar de volta par o front-end avisando que a ação é impossível.
		//PSEUDO-CÓDIGO: Interrompemos a função imediatamente, retornando um texto vazio e um aviso de erro formal.
		return "", errors.New("O avatar está sem roupas, não há o que desfazer")
	}

	//POR QUE: Em computação, as listas começam do índica 0. Se temos 3 roupas, elas estão nas posições 0, 1 e 2. Portanto, o topo é o tamanho (3) menos 1 = 2.
	//PSEUDO-CÓDIGO: Calculamos a posição(indice) do úlitimo item(-1) fazendo tamanho -1, para acessar o item mais recente adicionado.
	indiceTopo := tamanho - 1

	//POR QUE: Precisamos salvar qual roupa é essa na memória temporária antes de cortá-la do Slice, para podermos devolvê-la no final da função.
	//PSEUDO-CÓDIGO: Pegamos a roupa que está exatamente na posição 'indiceTopo' e copiamos para a variável "pecasRemovidas"
	pecaRemovida := p.pecasVestidas[indiceTopo]

	//POR QUE: O operador "[:indiceTopo]" diz ao Go para manter a 'janela' do slice apenas sobre os itens anteriores, efetivamente 'esquecendo' a última peça. Isso é a operação POP(desempilhar) da nossa Pilha, com custo O(1) | O(1):
	//PSEUDO-CÓDIGO: Atualizamos o Slice Original, cortando ele do início até ANTES do 'indiceTopo'
	p.pecasVestidas = p.pecasVestidas[:indiceTopo]

	//POR QUE: 'nil' em Go significa ausência de valor. Retornando erro nulo avisa ao sistema qe a operação foi um sucesso, e que o resultado é confiável.
	//PSEUDO-CÓDIGO: Finalizamos a func devolvendo o nome da peça removida e 'nil'(nulo) para o indicador de erro
	return pecaRemovida, nil
}

// POR QUE: O mundo externo (a func main) não tem acesso ao slice privado 'p.pecasVestidas'. Precisamos de um método seguro só para "ler" as roupas sem alterá-las.
// PSEUDO-CÓDIGO: Criamos o método VerLook para listar o que está no avatar.
func (p *ProvadorVirtual) VerLook() {

	// POR QUE: Melhora a experiência do usuário (UX). É estranho imprimir um título "Seu Look" e não mostrar nada abaixo dele.
	// PSEUDO-CÓDIGO: Verifica se o provador está vazio.
	tamanho := len(p.pecasVestidas)
	if tamanho == 0 {
		fmt.Println("\n O avatar está apenas de roupas íntimas (vazio).")

		// POR QUE: O 'return' solto no meio da função faz ela parar de executar aqui e voltar para o menu.
		// PSEUDO-CÓDIGO: Interrompe a execução do método.
		return
	}

	fmt.Println("\n--- LOOK ATUAL (Do topo para a base) ---")

	// POR QUE: Numa Pilha (LIFO), a última coisa que você veste é a que fica visível por cima de todas. O loop reverso garante que mostremos o look na ordem certa visualmente.
	// PSEUDO-CÓDIGO: Iniciamos um loop (for) que começa no último item do slice (tamanho - 1) e vai descendo até chegar no zero (i--).
	for i := tamanho - 1; i >= 0; i-- {
		fmt.Printf("- %s\n", p.pecasVestidas[i])
	}
	fmt.Println("-------------------------------------------")
}

// ___PASSO 3. O ProvadorVirtual em ação (execução - Interface Interativa)___

// PSEUDO-CÓDIGO: A func main é o ponto de partida. Quando ligamos o servidor, a CPU enra diretamente aqui para começar a executar o código. Tudo que estiver dentro da main() vai acontecer em sequência, como um roteiro de filme.
func main() {
	provador := ProvadorVirtual{}

	// POR QUE: A função simples do Go corta o texto se você digitar um espaço (ex: "Jaqueta Jeans" vira só "Jaqueta"). O bufio.Scanner cria um "radar" potente que lê a frase inteira com espaços.
	// PSEUDO-CÓDIGO: Criamos um Scanner conectado ao teclado do computador (os.Stdin).
	scanner := bufio.NewScanner(os.Stdin)

	// POR QUE: O programa não pode executar uma vez e morrer. Ele precisa ficar vivo, atendendo a cliente infinitamente, até ela decidir ir embora.
	// PSEUDO-CÓDIGO: Iniciamos um loop infinito (um 'for' sem condições de parada).
	for {
		fmt.Println("\n====== PROVADOR VIRTUAL ======")
		fmt.Println("1. Adicionar uma peça")
		fmt.Println("2. Desfazer última peça")
		fmt.Println("3. Ver Look Completo")
		fmt.Println("4. Sair do Provador")
		fmt.Print("Escolha uma opção: ")

		// POR QUE: Se não pausarmos aqui, o loop infinito vai rodar milhões de vezes por segundo e travar seu PC. Ele congela e só destrava quando você aperta o botão ENTER.
		// PSEUDO-CÓDIGO: O Scanner pausa o programa e escuta o teclado.
		scanner.Scan()

		// POR QUE: Precisamos limpar a sujeira. Se a cliente esbarrar na barra de espaço e digitar " 1 ", o 'TrimSpace' limpa as bordas e transforma em "1".
		// PSEUDO-CÓDIGO: Pegamos o texto digitado e removemos os espaços invisíveis ao redor dele.
		opcao := strings.TrimSpace(scanner.Text())

		// POR QUE: O 'switch' roteia a execução. É muito mais legível que escrever cinco blocos de 'if / else' seguidos para saber qual botão foi apertado.
		// PSEUDO-CÓDIGO: Analisamos o texto da 'opcao' escolhida.
		switch opcao {

		case "1":
			fmt.Print("Digite o nome da peça (ex: Casaco de Lã): ")

			// POR QUE: Pausamos o programa de novo, agora para ouvir qual roupa a cliente quer vestir.
			// PSEUDO-CÓDIGO: Escaneia o teclado novamente.
			scanner.Scan()
			novaPeca := strings.TrimSpace(scanner.Text())

			// POR QUE: Impede que a cliente adicione uma roupa "fantasma" sem nome no banco de dados.
			// PSEUDO-CÓDIGO: Verifica se o texto digitado não está vazio.
			if novaPeca != "" {

				// POR QUE: Como nosso método agora tem a trava de 5 peças e retorna erro, nós precisamos capturar esse possível erro.
				// PSEUDO-CÓDIGO: Executamos o AdicionarPeca e guardamos a resposta na variável 'err'.
				err := provador.AdicionarPeca(novaPeca)

				// POR QUE: Se 'err' for diferente de nulo, a cliente atingiu o limite de 5 peças. Mostramos o aviso.
				// PSEUDO-CÓDIGO: Tratamento do erro da trava de segurança.
				if err != nil {
					fmt.Printf("\nErro: %s\n", err)
				}
			} else {
				fmt.Println("\nErro: O nome da peça não pode ser vazio.")
			}

		case "2":
			pecaRemovida, err := provador.DesfazerLook()
			if err != nil {
				fmt.Printf("\nErro: %s\n", err)
			} else {
				fmt.Printf("\nAção desfeita! Você tirou: %s\n", pecaRemovida)
			}

		case "3":
			provador.VerLook()

		case "4":
			// POR QUE: O 'return' destroi a função main() imediatamente. Quando a main() morre, o loop é quebrado e o programa do terminal finaliza com sucesso.
			// PSEUDO-CÓDIGO: Despede-se da cliente e encerra o sistema executável.
			fmt.Println("\nSaindo do provador... Volte sempre!")
			return

		default:
			// POR QUE: Rota de escape. Se a cliente digitar "9" ou "batata", o sistema não entende, mas também não quebra. Ele cai aqui e o loop reinicia o menu.
			// PSEUDO-CÓDIGO: Captura qualquer entrada que não seja de 1 a 4.
			fmt.Println("\nOpção inválida. Digite um número de 1 a 4.")
		}
	}
}

//POR QUE:
//PSEUDO-CÓDIGO:
