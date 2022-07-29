import React from 'react';

const Header = () =>{
    const logoTitle = 'My First React App';
    return (
        <header className='header row mt-3'>
            <div className='logo col-4'>
                <a href="#">{logoTitle}</a>
            </div>
            <div className='col-4'></div>
            <div className='col-4 search'>
                <div className='input-group'>
                    <input type="text" className="form-control" placeholder="Search Keyword" />
                    <div className='input-group-append'>
                        <button className='btn btn-outline-secondary'>Search </button>
                    </div>
                </div>
            </div>
        </header>
    )
}
export default Header;