
const template = `
        <style>
        .navbar {
            position: fixed;
            top: 0;
            left: 0;
            width: 4rem;
            height: 100%;
            background: linear-gradient(90deg, hsl(0, 12%, 16%), hsl(0, 13%, 26%), hsl(0, 16%, 19%));
            transition: ease 0.3s;

            &:hover {
                width: 12rem;
            }

            &:hover .navegacao {
                display: flex;
            }

            &:hover .icone {
                width: 4rem;
                height: 4rem;
                padding: 2rem;
                margin-bottom: 2rem;
                margin-left: 3rem;
            }
        }

        .icone {
            width: 3rem;
            height: 3rem;
            padding: 1.5rem;
            margin-bottom: 1.5rem;
            transition: ease 0.3s;
        }

        .navegacao {
            display: none;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            width: 100%;
        }

        .navegacao a {
            width: 100%;
            height: 4rem;
            color: whitesmoke;
            text-decoration: none;
            font-size: 1.2rem;
            display: flex;
            align-items: center;

            &:hover {
                cursor: pointer;
                background-color: hsl(0, 11%, 39%);
            }
        }
    </style>
    <nav class="navbar">
        <div class="icone">
            <img src="" alt="ICONE">
        </div>
        <div class="navegacao">
            <a href="./inicioPagina.html">Início</a>
            <a href="../formularioPaciente/FormularioPaciente.html">Configurações</a>
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
