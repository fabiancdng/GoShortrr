import { Heading, Link } from '@chakra-ui/layout';
import { Table, Td, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/table';
import React, { useEffect, useState } from 'react';
import { host, getShortlinkList } from '../../adapters/ShortlinkAdapter';

const Shortlinks = () => {
    const [shortlinks, setShortlinks] = useState([]);

    useEffect(() => {
        getShortlinkList()
            .then(shortlinkList => {
                setShortlinks(shortlinkList);
            })
            // TODO: Handle error
            .catch(err => {});
    }, []);

    return (
        <>
            <Heading width='90%'>Your shortlinks</Heading>
            <Table variant="simple" width='90%' mt={7}>
                <Thead>
                    <Tr>
                        <Th>Shortlink</Th>
                        <Th>Redirects to</Th>
                        <Th>Created</Th>
                        <Th>Actions</Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {shortlinks.map((shortlink) => (
                        <Tr key={ shortlink.id }>
                            <Td>
                                <Link href={ host + shortlink.short } target='goshortrr'>
                                    { host + shortlink.short }
                                </Link>
                            </Td>
                            <Td>
                                <Link href={ shortlink.link } target='goshortrr'>
                                    { shortlink.link }
                                </Link>
                            </Td>
                            <Td>{ shortlink.created.replace('T', ' - ').replace('Z', '') }</Td>
                            <Td>Coming soon!</Td>
                        </Tr>
                    ))}
                </Tbody>
            </Table>
        </>
    );
}

export default Shortlinks;
