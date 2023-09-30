import React, { useEffect, useState } from 'react';
import { getAllCarriers } from '../api/carrier';
import '../css/carrier.css';

const CarriersForm: React.FC = () => {
    const [carriers, setCarriers] = useState<any[]>([]);

    useEffect(() => {
        const fetchCarriers = async () => {
            try {
                const allCarriers = await getAllCarriers();
                setCarriers(allCarriers);
            } catch (error) {
                console.error('Error fetching carriers:', error);
                // Handle error
            }
        };

        fetchCarriers();
    }, []);

    return (
        <div>
            <h1>List of Carriers</h1>
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Carrier</th>
                        <th>Phone</th>
                    </tr>
                </thead>
                <tbody>
                        {carriers.map((carrier: any) => (
                            <tr>
                            <td>{carrier.id}</td>
                            <td>{carrier.name}</td>
                            <td>{carrier.telephone}</td>
                            </tr>
                        ))}
                </tbody>
            </table>
        </div>
    );
};

export default CarriersForm;