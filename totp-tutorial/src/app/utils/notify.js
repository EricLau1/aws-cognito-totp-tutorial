import { toast } from 'react-toastify';

const defaultStyle = {
    position: "top-right",
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
}

function Notify(message) {
    return {
        asSuccess: () => toast.success(message, defaultStyle),
        asError: () => toast.error(message, defaultStyle),
        asInfo: () => toast.info(message, defaultStyle),
        asWarning:() => toast.warn(message, defaultStyle),
    };
};

export default Notify;