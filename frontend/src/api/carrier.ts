import axios from 'axios';

//const baseUrl = process.env.PUBLIC_URL;
const apiPort = process.env.REACT_APP_API_SERVER_PORT;

const apiUrl = `${window.location.protocol}//${window.location.hostname}:${apiPort}/api/carrier/all`;
//const apiUrl = `${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/carrier/all`;
//const apiUrl = `/api/carrier/all`;

export const getAllCarriers = async () => {
    try {
        console.log('API URL:', apiUrl); // Add this line
        const response = await axios.get(apiUrl);
        //const response = await axios.get(`http://localhost:3030/api/carrier/all`);
        return response.data;
    } catch (error) {
        console.error('Error retrieving carriers:', error);
        throw error;
    }
};
