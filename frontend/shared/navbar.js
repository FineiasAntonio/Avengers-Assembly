
const template = `
        <style>
        @import "./estilos-globais.css";

        .navbar {
            position: fixed;
            z-index: 100;
            top: 0;
            left: 0;
            width: 4rem;
            height: 100%;
            background: var(--cor-primaria);
            transition: ease 0.3s;
            font-family: "Inria Serif", "Times New Roman";

            &:hover {
                width: 10rem;
            }

            &:hover .navegacao {
                display: flex;
                opacity: 1;
                transition:
                    opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            }

            &:hover .icone {
                width: 4rem;
                height: 4rem;
                padding: 1rem;
                margin-bottom: 2rem;
                margin-left: 3rem;
            }
        }

        .icone{
            width: 2rem;
            height: 2rem;
            padding: 1.5rem;
            margin-bottom: 1.5rem;
            transition: ease 0.3s;
        }

        .navegacao {
            flex-direction: column;
            align-items: center;
            justify-content: center;
            width: 100%;
            opacity: 0;
            transition:
                opacity 0.10s cubic-bezier(0.4, 0, 0.2, 1);
        }

        .navegacao a {            
            width: 100%;
            height: 3rem;
            color: whitesmoke;
            text-decoration: none;
            font-size: 1rem;
            display: flex;
            align-items: center;
            padding: 1rem;

            &:hover {
                cursor: pointer;
                background-color: var(--cor-quinaria);
            }
        }
        
        @media (min-width: 1366px) {
            .navbar {
                width: 6rem;

                &:hover {
                    width: 12rem;
                }
            }

            .navegacao a {
                font-size: 1.3rem;
                height: 4rem;
            }
        }

        @media (max-width: 1000px) {
            .navbar {
                width: 2rem;

                &:hover {
                    width: 12rem;
                }
            } 
        }

        @media (max-width: 700px) {
            .navbar {
                &:hover {
                    width: 10rem;
                }
            }
        }

        @media (max-width: 560px) {
            .navbar {
                width: 1.5rem;
            }
        }

    </style>
    <nav class="navbar">
        <div class="icone">
            <img src="../../assets/barIcon.png" alt="ICONE">
        </div>
        <div class="navegacao">
            <a href="">Início</a>
            <a href="">Configurações</a>
            <a href="">Ajuda</a>
            <a href="">Sair</a>
        </div>        
    </nav>
    `
class Navbar extends HTMLElement {
  connectedCallback() {
    this.innerHTML = template;
  }
}

customElements.define("custom-navbar", Navbar);
