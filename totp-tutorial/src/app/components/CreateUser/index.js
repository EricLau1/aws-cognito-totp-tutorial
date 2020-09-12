import React from 'react';
import axios from 'axios';
import Notify from '../../utils/notify';

const defaultForm = { email: '', password: '', created: undefined };

class CreateUser extends React.Component {
    constructor(props) {
        super(props);
        this.state = { ...defaultForm };
    }
    render() {
        return (
            <div>
                <h2>Criar Conta</h2>
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
                <button type="button" onClick={() => this.create()}>
                    Criar
                </button>
                <h2>Confirmar Conta</h2>
                <ConfirmUser email={this.state.created} />
            </div>
        )
    }

    async create() {
        const { email, password } = this.state;

        try {
            const data = await axios.post('http://localhost:8080/api/totp/signup', { email, password });
            console.log(data);
            this.setState({ ...defaultForm, created: email });
            Notify('Conta criada com sucesso!').asSuccess();
        } catch(err) {
            console.info(err.response);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }
    }
}

const defaultConfirmForm = { email: '', code: '' };

class ConfirmUser extends React.Component {
    constructor(props){
        super(props);
        this.state = { ...defaultConfirmForm };
    }
    render() {
        return (
            <div>
                <div>
                    <label>Email:</label>
                    <input 
                        name="email"
                        onChange={e => this.setState({email: e.target.value})}
                        placeholder="Email"
                        value={this.props.email || this.state.email}
                    />
                </div>
                <div>
                    <label>Código:</label>
                    <input 
                        name="code"
                        onChange={e => this.setState({code: e.target.value})}
                        placeholder="Código"
                        value={this.state.code}
                    />
                </div>
                <button type="button" onClick={() => this.confirm()}>
                    Confirmar
                </button>
            </div>
        );
    }

    async confirm() {
        const { email, code } = this.state;

        let values = {
            email: email !== '' ? email : this.props.email,
            code
        };

        try {
            const data = await axios.post('http://localhost:8080/api/totp/confirm-signup', values);
            console.log({data});
            Notify('Conta confirmada com sucesso!').asSuccess();
            this.setState({code: '', email: ''});
        } catch(err) {
            console.error(err);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }
    }

}

export default CreateUser;