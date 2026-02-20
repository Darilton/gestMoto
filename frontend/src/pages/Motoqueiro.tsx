import { useEffect, useState } from 'react';
import { fetchData } from '../services/api';
import type { IMotoqueiro } from '../models/motoqueiro';

function Motoqueiro() {
    const [data, setData] = useState<IMotoqueiro[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        fetchData()
            .then(data => {
                setData(data["motoqueiros"]);
                setLoading(false);
            })
            .catch(err => {
                setError(err);
                setLoading(false);
            });
    }, []);

    const listItems = data.map(person =>
        <div key={person.id} className='h-20 w-full border rounded-sm'>
            <h1>{person.nome}</h1>
            <h1>{person.telefone}</h1>
        </div>
    );

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error loading data!</p>;
    return (
        <>
        <div className="flex flex-col p-2 gap-2">
            <h1>Motoqueiros</h1>
            {listItems}
        </div>
        </>
    );
}

export default Motoqueiro;