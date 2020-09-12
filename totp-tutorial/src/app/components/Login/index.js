import React from 'react';
import axios from 'axios';
import Notify from '../../utils/notify';
import Mfa from '../Mfa';

const defaultLoginForm = { email: '', password: '', authInfo: undefined, authData: {} };

class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = { ...defaultLoginForm };
    }
    render() {
        return (
            <div>
                <h2>Login</h2>
                <div>
                    <div>
                        <label>Email:</label>
                        <input 
                            name="email"
                            onChange={e => this.setState({email: e.target.value})}
                            placeholder="Email"
                            value={this.state.email}
                        />
                    </div>
                    <div>
                        <label>Senha:</label>
                        <input 
                            name="senha"
                            onChange={e => this.setState({password: e.target.value})}
                            placeholder="Senha"
                            value={this.state.password}
                        />
                    </div>
                    <button type="button" onClick={() => this.login()}>
                        Entrar
                    </button>
                </div>
                <Auth auth={this.state.authInfo} data={this.state.authData} confirmChallenge={this.confirmChallenge.bind(this)} />
            </div>
        );
    }

    async login() {
        const { email, password } = this.state;

        try {
            const data = await axios.post('http://localhost:8080/api/totp/login', { email, password });
            console.log(data);
            this.setState({ ...defaultLoginForm, authInfo: data.data });
            this.getUser(email);
        } catch(err) {
            console.info(err.response);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }
    }

    async getUser(email) {
        try {
            const response = await axios.get('http://localhost:8080/api/totp/user/' + email);
            console.log(response);
            this.setState({ authData: response.data });
        } catch(err) {
            console.info(err.response);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }
    }

    async confirmChallenge(token, username, auth_info) {
        try {
            const response = await axios.post('http://localhost:8080/api/totp/confirm-login', { username, token, auth_info });
            console.log('CHALLENGE_COMPLETED', {response});
            Notify('Desafio completado com sucesso!').asSuccess();
            this.getUser(username);
            this.setState({authInfo: response.data});
        } catch(err) {
            console.info(err.response);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }
    }
}

class Auth extends React.Component {
    render() {
        const { auth, data } = this.props;
        if(!auth || !this.isLogged(auth.AuthenticationResult)) {

            if(auth && auth.ChallengeName) {
                return (
                    <div>
                        <h2>Completar Desafio</h2>
                        <pre>{JSON.stringify(auth || {}, null, 2)}</pre>
                        <RespondChallenge authInfo={auth} username={data.Username} confirm={this.props.confirmChallenge} />
                    </div>
                );
            }

            return <p>Offline...</p>
        }

        console.log(data);

        return (
            <div>
                <h2>MFA</h2>
                <div>
                    <Mfa accessToken={auth.AuthenticationResult.AccessToken || ''} username={data.Username} />
                </div>
                <h2>Auth Info</h2>
                <div>
                    <pre>{JSON.stringify(data || {}, null, 2)}</pre>
                </div>
            </div>
        );
    }

    isLogged(auth) {
        return auth &&
               auth.AccessToken && auth.AccessToken !== '' &&
               auth.IdToken && auth.IdToken !== '' &&
               auth.RefreshToken && auth.RefreshToken;
    }
}

class RespondChallenge extends React.Component {
    constructor(props) {
        super(props);
        this.state = { token: '' }
    }
    render() {
        return (
            <div>
                <div>   
                    <label>Token</label>
                    <input 
                        name="token"
                        type="text"
                        onChange={e => this.setState({ token: e.target.value })}
                        placeholder="Token"
                        value={this.state.code}
                    />
                    <button type="button" onClick={() => this.confirm()}>Confirmar</button>
                </div>
            </div>
        );
    }

    async confirm() {
        const { username, authInfo } = this.props;
        const { token } = this.state;
        this.props.confirm(token, username, authInfo);
    }
}

export default Login;