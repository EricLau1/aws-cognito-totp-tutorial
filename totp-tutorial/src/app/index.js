import React from 'react';
import CreateUser from './components/CreateUser';
import Login from './components/Login';
import { ToastContainer } from 'react-toastify';

class App extends React.Component {
    render() {
        return (
            <div>
                <ToastContainer />
                <h1>TOTP Tutorial</h1>
                <hr />
                <CreateUser />
                <hr />
                <Login />
            </div>
        );
    }
}

export default App;