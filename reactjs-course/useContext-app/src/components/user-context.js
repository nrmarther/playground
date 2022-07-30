import React from'react';

const user = {
    name: 'Nathan',
    job: 'React Developer'
}

export const userContext = React.createContext(user)