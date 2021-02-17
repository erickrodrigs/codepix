import React from 'react';

type Props = {

};

export default function BankAccounts(props: Props) {
  return (
    <div>
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="container-fluid">
          <a className="navbar-brand" href="#">
            <img src="/img/icon_banco.png" alt="Bank"/>
            <div>
              <span>Cod - 001</span>
              <h2>BBX</h2>
            </div>
          </a>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarSupportedContent">
            <ul className="navbar-nav me-auto mb-2 mb-lg-0">
              <li className="nav-item">
                <img src="/img/icon_user.png" alt="User"/>
                <p>
                  Proprietário | C/C: Número da conta
                </p>
              </li>
            </ul>
          </div>
        </div>
      </nav>
      <main>
        <div className="container">
          <h1>Contas bancárias</h1>
        </div>
      </main>
      <footer>
        <img src="/img/logo_pix.png" alt="Codepix"/>
      </footer>
    </div>
  )
}
