import React from 'react';
import axios from 'axios';
import Notify from '../../utils/notify';
import QRCode from 'qrcode.react';

class Mfa extends React.Component {
    constructor(props) {
        super(props);
        this.state = { secretCode: '', code: '' };
    }

    render() {
        return (
            <div>
                {
                    this.state.secretCode !== '' ? (
                        <div>
                            <div>   
                                <label>Código</label>
                                <input 
                                    name="code"
                                    type="text"
                                    onChange={e => this.setState({ code: e.target.value })}
                                    placeholder="Código"
                                />
                                <button type="button" onClick={() => this.activate()}>Ativar</button>
                            </div>
                            <div>
                                <p><strong>Escanear Código</strong></p>
                                <QRCode value={this.buildUrl(this.state.secretCode, this.props.username)} />
                            </div>
                        </div>
                    ) : (
                        <div>
                            <button type="button" onClick={() => this.associate()}>
                                Associar Token
                            </button>
                        </div>
                    )
                }
            </div>
        );
    }

    buildUrl(secretCode, username) {
        return "otpauth://totp/tutorial:" + username + "?secret=" + secretCode + "&issuer=tutorial";
    }

    async associate() {
        const { accessToken } = this.props;

        try {
            const data = await axios.post('http://localhost:8080/api/totp/associate', { accessToken });
            console.log({data});
            Notify('Token associado com sucesso!').asSuccess();
            this.setState({secretCode: data.data.SecretCode});
        } catch(err) {
            console.error(err);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }

    }

    async activate() {
        const { accessToken } = this.props;

        try {
            const data = await axios.post('http://localhost:8080/api/totp/confirm', { accessToken, code: this.state.code });
            console.log({data});
            Notify('Totp ativado com successo!').asSuccess();
            this.setState({secretCode: data.data.SecretCode});
        } catch(err) {
            console.error(err);
            Notify(err.response.data.error || 'Ocorreu um erro! :(').asError();
        }

    }
}

export default Mfa;