package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ProvadorVirtual struct {
	pecasVestidas []string
}

func (p *ProvadorVirtual) AdicionarPeca(peca string) error {

	tamanho := len(p.pecasVestidas)

	if tamanho >= 5 {

		return errors.New("limite de peças atingido! O provador suporta no máximo 5 peças de roupa")
	}

	p.pecasVestidas = append(p.pecasVestidas, peca)
	fmt.Printf("Sucesso! Você vestiu: %s\n", peca)

	return nil
}

func (p *ProvadorVirtual) DesfazerLook() (string, error) {

	tamanho := len(p.pecasVestidas)

	if tamanho == 0 {

		return "", errors.New("o avatar está sem roupas, não há o que desfazer")
	}

	indiceTopo := tamanho - 1

	pecaRemovida := p.pecasVestidas[indiceTopo]

	p.pecasVestidas = p.pecasVestidas[:indiceTopo]

	return pecaRemovida, nil
}

func (p *ProvadorVirtual) VerLook() {

	tamanho := len(p.pecasVestidas)
	if tamanho == 0 {
		fmt.Println("\n O avatar está apenas de roupas íntimas (vazio).")

		return
	}

	fmt.Println("\n--- LOOK ATUAL (Do topo para a base) ---")

	for i := tamanho - 1; i >= 0; i-- {
		fmt.Printf("- %s\n", p.pecasVestidas[i])
	}
	fmt.Println("-------------------------------------------")
}

func main() {
	provador := ProvadorVirtual{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n====== PROVADOR VIRTUAL ======")
		fmt.Println("1. Adicionar uma peça")
		fmt.Println("2. Desfazer última peça")
		fmt.Println("3. Ver Look Completo")
		fmt.Println("4. Sair do Provador")
		fmt.Print("👉🏾 Escolha uma opção: ")

		scanner.Scan()

		opcao := strings.TrimSpace(scanner.Text())

		switch opcao {

		case "1":
			fmt.Print("Digite o nome da peça (ex: Casaco de Lã): ")

			scanner.Scan()
			novaPeca := strings.TrimSpace(scanner.Text())

			if novaPeca != "" {

				err := provador.AdicionarPeca(novaPeca)

				if err != nil {
					fmt.Printf("\n❌ Erro: %s\n", err)
				}
			} else {
				fmt.Println("\n❌ Erro: O nome da peça não pode ser vazio.")
			}

		case "2":
			pecaRemovida, err := provador.DesfazerLook()
			if err != nil {
				fmt.Printf("\n❌ Erro: %s\n", err)
			} else {
				fmt.Printf("\n✂️ Ação desfeita! Você tirou: %s\n", pecaRemovida)
			}

		case "3":
			provador.VerLook()

		case "4":
			fmt.Println("\n👋🏾 Saindo do provador... Volte sempre!")
			return

		default:
			fmt.Println("\n❌ Opção inválida. Digite um número de 1 a 4.")
		}
	}
}
