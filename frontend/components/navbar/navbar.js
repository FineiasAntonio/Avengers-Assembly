import { deslogar } from "../../shared/gerenciador-login.js";

const template = `
        <style>
        @import "../../shared/estilos-globais.css";

        .navbar {
            display: flex;
            flex-direction: column;
            align-items: center;
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
                width: 100%;
                height: 4rem;
                padding: 1rem;
                margin-bottom: 2rem;
            }
        }

        .icone{
            display: flex;
            align-items: center;
            justify-content: center;
            width: 100%;
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
                width: 4rem;

                &:hover {
                    width: 10rem;
                }
            }

            .navegacao a {
                font-size: 1rem;
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

            .icone {
                display: hidden;
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
            <img src="../../assets/barIcon.png" alt="ICONE" class="iconeImg">
        </div>
        <div class="navegacao">
            <a href="../inicioPagina/inicioPagina.html">Início</a>
            <a href="../perfilUsuario/perfilUsuario.html">Perfil</a>
            <a href="../centralAjuda/centralAjuda.html">Ajuda</a>
            <a id="deslogar-opcao">Sair</a>
        </div>        
    </nav>
    `
class Navbar extends HTMLElement {
  async connectedCallback() {
    this.innerHTML = template;

    document.getElementById("deslogar-opcao").addEventListener('click', (e) => {
        e.preventDefault()
        deslogar()
    })
  }
}

customElements.define("custom-navbar", Navbar);
