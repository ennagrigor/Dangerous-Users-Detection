import React from 'react';
import {Jumbotron, Container} from 'reactstrap';
import './Header.css'

const Header = () => {
    return (
        <div>
            <Jumbotron fluid className="header">
                <Container fluid>
                    <h2>DANGEROUS USERS DETECTION APPLICATION</h2>
                    <img src={window.location.origin + '/logo.jpeg'} className="img-center" alt="logo"/>
                    <p>The purpose of the product is to scan social media platforms that have
                        public APIs and extract text via posts, tweets etc.
                        Then it will monitor and identify keywords according to a pre-defined "dangerous" vocabulary so
                        we can maintain a list of suspicious people who may pose a social danger.</p>
                </Container>
            </Jumbotron>
        </div>
    );
};

export default Header;
