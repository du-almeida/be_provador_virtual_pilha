// O PORQUE: O compilador do Go exige que um "package main" para saber que deve gerar um arquivo executável (.exe ou binário) a partir daqui, e não apenas uma bilbioteca de apoio
// PSEUDO-CÓDIGO: Declaramos que este arquivo é o coração do programa (o pacote principal sempre é o "Package Main")
package main

//O PORQUE: Não precisamos reinventar a roda. Trazemos pacotes que nos auxiliam a escrever e executar nosso código(nesse caso de errors e formatação de texto "fmt")
//PSEUDO-CÓDIGO: Importamos ferramentas externas prontas para usarmos no nosso código
import (
	"fmt"
)

// ___PASSO 1. Classe (O molde)___
// O PORQUE: Este é o "Molde de papel" no nosso ateliê. Ele define como o objeto será construído na memória RAM quando for instanciado(definido,criado)
// PSEUDO-CÓDIGO: Criamos um novo tipo de dado estruturado (struct) chamado de ProvadorVirtual
type provadorVirtual struct {

	//O PORQUE: A letra minúscula "p" em pecasProvadas é uma regra do Go para "ENCAPSULAMENTO". Ela proíbe que outros arquivos fora deste pacote toquem nesse Slice diretamente, obgirando o uso dos métodos seguros.
	//PSEUDO-CÓDIGO: Criamos um ATRIBUTO privado chamado "pecasProvadas", que é uma Lista (Slice) restrita para textos(strings)
	pecasProvadas []string
}

// ___PASSO 2. Os Métodos (O comportamento)__
// O PORQUE: O "(p *ProvadorVirtual)" é o Method Receiver(pertence a um tipo específico). O "*"(ponteiro) diz à CPU: "Não crie uma cópia do manequim. Vá até o endereço de memória do manequim real "p" e altere ele"
// PSEUDO-CÓDIGO: Criamos o método AdicionarPeca que recebe um texto (peca). Ele pertence exclusivamente ao ProvadorVirtual.
func (p *ProvadorVirtual) AdicionarPeca(peca string) {

	//O PORQUE: A função "append" é ótimizada pelo Go para alocar mais espaço na memória apenas se a nossa "caixa"(capacidade do slice) estiver cheia. Esta é a operação PUSH(empilhar) da nossa Pilha, com custo O(1) | O(1):
	//PSEUDO-CÓDIGO: Pegamos o slice atual, anexamos a uma nova "peca no final dele, e guardamos o resultado de volta no mesmo lugar.
	p.pecasProvadas = append(p.pecasProvadas, peca)
	fmt.Printf("Vestido: %s\n", peca)
}

// O PORQUE: O retorno multiplo "(string, error)" é o padrão ouro da segurança no Go Se algo der errado, ele avisa, se der certo ele entrega os dados.
// PSEUDO-CÓDIGO: Criamos o método DesfazerLook. Ele não recebe parâmetros, mas devolve Duas coisas: um texto( a peça) e um erro.
func (p *ProvadorVirtual) DesfazerLook(string, error) {

	//O PORQUE: Precisamos saber a quantidade exata para descobrir matematicamente quem
	//PSEUDO-CÓDIGO: Medimos(len) quantos itens existem dentro do slice "pecasProvadas"
	tamanho := len(p.pecasProvadas)

}

//O PORQUE:
//PSEUDO-CÓDIGO:
