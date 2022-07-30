import React, {useContext} from 'react';
import { colorContext } from './color-context';
import { userContext } from './user-context';

const SampleContext = () =>{
    const colors = useContext(colorContext);
    const user = useContext(userContext)
    return (
        <>
            <div style={{backgroundColor: colors.red}}>This is Red</div>
            <div style={{backgroundColor: colors.blue}}>This is Blue</div>
            <div style={{backgroundColor: colors.purple}}>This is Purple</div>
            <div style={{backgroundColor: colors.green}}>This is Green</div>
            <br />
            <div>Name: {user.name}</div>
            <div>Job: {user.job}</div>


        </>
    )
}
export default SampleContext;