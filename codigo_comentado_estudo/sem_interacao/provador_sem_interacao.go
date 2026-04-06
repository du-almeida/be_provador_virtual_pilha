// PSEUDO-CÓDIGO: Declaramos que este arquivo é o coração do programa (o pacote principal).
// POR QUE: O compilador do Go exige um "package main" para saber que deve gerar um arquivo executável (.exe ou binário) a partir daqui, e não apenas uma biblioteca de apoio.
package main

// PSEUDO-CÓDIGO: Importamos ferramentas externas prontas para usarmos no nosso código.
// POR QUE: Não precisamos reinventar a roda. Trazemos o pacote de erros (errors) e o de formatação de texto (fmt).
import (
	"errors"
	"fmt"
)

// --- 1. A CLASSE (O Molde) ---

// PSEUDO-CÓDIGO: Criamos um novo tipo de dado estruturado (struct) chamado ProvadorVirtual.
// POR QUE: Este é o "Molde de Papel" no nosso ateliê. Ele define como o objeto será construído na memória RAM quando for instanciado.
type ProvadorVirtual struct {

	// PSEUDO-CÓDIGO: Criamos um Atributo privado chamado 'pecasVestidas', que é uma Lista (Slice) restrita para textos (strings).
	// POR QUE: A letra minúscula 'p' em pecasVestidas é uma regra do Go para "Encapsulamento". Ela proíbe que outros arquivos fora deste pacote toquem neste Slice diretamente, obrigando o uso dos métodos seguros.
	pecasVestidas []string
}

// --- 2. OS MÉTODOS (O Comportamento) ---

// PSEUDO-CÓDIGO: Criamos o método AdicionarPeca que recebe um texto (peca). Ele pertence exclusivamente ao ProvadorVirtual.
// POR QUE: O `(p *ProvadorVirtual)` é o Method Receiver. O `*` (ponteiro) diz à CPU: "Não crie uma cópia do manequim. Vá até o endereço de memória do manequim real 'p' e altere ele".
func (p *ProvadorVirtual) AdicionarPeca(peca string) {

	// PSEUDO-CÓDIGO: Pegamos o slice atual, anexamos a nova 'peca' no final dele, e guardamos o resultado de volta no mesmo lugar.
	// POR QUE: A função `append` é otimizada pelo Go para alocar mais espaço na memória apenas se a nossa "caixa" (capacidade do slice) estiver cheia. Esta é a operação PUSH (Empilhar) da nossa Pilha, com custo O(1).
	p.pecasVestidas = append(p.pecasVestidas, peca)

	// PSEUDO-CÓDIGO: Imprimimos na tela (console) o que acabou de ser vestido.
	// POR QUE: Feedback visual para sabermos que a operação foi concluída no servidor.
	fmt.Printf(" Vestindo: %s\n", peca)
}

// PSEUDO-CÓDIGO: Criamos o método DesfazerLook. Ele não recebe parâmetros, mas devolve DUAS coisas: um texto (a peça) e um erro.
// POR QUE: O retorno múltiplo `(string, error)` é o padrão ouro de segurança no Go. Se algo der errado, ele avisa; se der certo, ele entrega o dado.
func (p *ProvadorVirtual) DesfazerLook() (string, error) {

	// PSEUDO-CÓDIGO: Medimos quantos itens existem dentro do slice 'pecasVestidas' agora e guardamos na variável 'tamanho'.
	// POR QUE: Precisamos saber a quantidade exata para descobrir matematicamente quem é o último da fila (o topo da pilha).
	tamanho := len(p.pecasVestidas)

	// PSEUDO-CÓDIGO: Verificamos se o tamanho é exatamente igual a zero (vazio).
	// POR QUE: Isso é "Programação Defensiva". Se tentarmos tirar uma peça de um slice vazio, o Go dá um erro fatal de memória (Panic) e derruba o servidor.
	if tamanho == 0 {

		// PSEUDO-CÓDIGO: Interrompemos a função imediatamente, retornando um texto vazio e um aviso de erro formal.
		// POR QUE: O `errors.New` cria o objeto de erro que o Back-end vai enviar de volta para o Front-end avisando que a ação é impossível.
		return "", errors.New("o avatar já está sem roupas, não há o que desfazer")
	}

	// PSEUDO-CÓDIGO: Calculamos a posição (índice) do último item fazendo Tamanho - 1.
	// POR QUE: Em computação, as listas começam no índice 0. Se temos 3 roupas, elas estão nas posições 0, 1 e 2. Portanto, o topo é o tamanho (3) menos 1 = 2.
	indiceTopo := tamanho - 1

	// PSEUDO-CÓDIGO: Pegamos a roupa que está exatamente na posição 'indiceTopo' e copiamos para a variável 'pecaRemovida'.
	// POR QUE: Precisamos salvar qual roupa é essa na memória temporária antes de cortá-la do Slice, para podermos devolvê-la no final da função.
	pecaRemovida := p.pecasVestidas[indiceTopo]

	// PSEUDO-CÓDIGO: Atualizamos o Slice original, cortando ele do início até ANTES do 'indiceTopo'.
	// POR QUE: O operador `[:indiceTopo]` diz ao Go para manter a "janela" do slice apenas sobre os itens anteriores, efetivamente "esquecendo" a última peça. Esta é a operação POP (Desempilhar).
	p.pecasVestidas = p.pecasVestidas[:indiceTopo]

	// PSEUDO-CÓDIGO: Finalizamos a função devolvendo o nome da peça removida e 'nil' (nulo) para o erro.
	// POR QUE: `nil` em Go significa ausência de valor. Retornar erro nulo avisa ao sistema que a operação foi um sucesso absoluto.
	return pecaRemovida, nil
}

// --- 3. O OBJETO NO MUNDO REAL (Execução) ---

// PSEUDO-CÓDIGO: A função main é o ponto de partida. Quando ligamos o servidor, a CPU entra diretamente aqui.
func main() {

	// PSEUDO-CÓDIGO: Criamos uma variável 'provador' e instanciamos nela uma cópia real do ProvadorVirtual.
	// POR QUE: Neste momento, o computador vai até a memória Heap e aloca um espaço físico para guardar as roupas daquela cliente específica (O Objeto nasceu).
	provador := ProvadorVirtual{}

	fmt.Println("--- INICIANDO PROVADOR ---")

	// PSEUDO-CÓDIGO: Usamos o objeto instanciado e chamamos o método de adicionar.
	// POR QUE: Ao fazer 'provador.', o Go secretamente passa o endereço de memória deste provador para dentro da variável 'p' lá no método.
	provador.AdicionarPeca("Vestido de Algodão")
	provador.AdicionarPeca("Jaqueta Jeans")
	provador.AdicionarPeca("Colar de Prata")

	fmt.Println("\n--- CLIENTE CLICOU EM DESFAZER ---")

	// PSEUDO-CÓDIGO: Chamamos a função de desfazer e preparamos duas variáveis ('pecaTirada' e 'err') para receber as duas respostas dela.
	// POR QUE: Toda função em Go que pode dar problema te obriga a capturar a variável de erro. É chato, mas é o que torna o Go a linguagem mais estável do mercado.
	pecaTirada, err := provador.DesfazerLook()

	// PSEUDO-CÓDIGO: Verificamos se a variável de erro é diferente (!=) de nulo.
	// POR QUE: Se for diferente de nulo, significa que a função encontrou um problema (ex: o slice estava vazio). Tratamos o erro aqui em vez de deixar o programa explodir.
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		// PSEUDO-CÓDIGO: Se o erro for nulo, a operação deu certo, então mostramos o resultado na tela.
		fmt.Printf("Desfez a última ação. Peça removida: %s\n", pecaTirada)
	}

	// PSEUDO-CÓDIGO: Chamamos o desfazer de novo, mas usamos o '_' (underline) no lugar do erro.
	// POR QUE: O `_` (Blank Identifier) diz ao compilador: "Eu sei que essa função me devolve um erro, mas eu vou ignorá-lo conscientemente desta vez, jogue-o no lixo".
	pecaTiradaDois, _ := provador.DesfazerLook()
	fmt.Printf("Desfez novamente. Peça removida: %s\n", pecaTiradaDois)
}
