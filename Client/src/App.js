import React, {Component} from 'react';
import Users from "./Users";
import Header from "./Header";
import "./App.css"
import Tweets from "./Tweets";

class App extends Component {
    render() {
        return (
            <div className="app-font">
                <Header/>
                <Users/>
                <Tweets/>
            </div>
        );
    }
}

export default App;
