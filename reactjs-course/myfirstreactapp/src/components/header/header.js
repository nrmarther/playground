import React from 'react';

const Header = () =>{
    let headerClass='header';
    const logoTitle = 'My First React App';
    return (
        <header className={headerClass}>
            <div className='logo'>
                <a href="#">{logoTitle}</a>
            </div>
            header here
        </header>
    )
}
export default Header;