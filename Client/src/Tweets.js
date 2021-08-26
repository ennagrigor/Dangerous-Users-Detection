import React, {Component} from 'react';
import {InputGroup, InputGroupAddon, Button, Input, Table} from 'reactstrap';
import "./Tweets.css"
import axios from "axios";

class Tweets extends Component {
    constructor(props) {
        super(props);
        this.state = {
            searchText: "",
            tweets: []
        };
    }

    componentDidMount() {
        this.getDangerousTweets()
    }

    handleOnSearchChange = (e) => {
        const searchText = e?.target?.value;
        this.setState({searchText});
    };

    handleSearchTweets = () => {
        const {searchText} = this.state;
        const searchSize = 1000;
        this.getTweets(searchSize, searchText);
    };

    getTweets = (size, query) => {
        axios.post("http://localhost:8080/tweets", {size: size, query: query})
            .then(response => {
                const tweets = response?.data
                this.setState({tweets});
            })
    }

    getDangerousTweets = () => {
        this.getTweets(10, "*")
    }

    render() {
        const {tweets} = this.state
        const header = <h2 className="header">DANGEROUS TWEETS</h2>;

        return (
            <div className="tweets">
                {header}
                <InputGroup>
                    <Input onChange={this.handleOnSearchChange}/>
                    <InputGroupAddon addonType="append">
                        <Button color="secondary" onClick={this.handleSearchTweets}>Search Tweets</Button>
                    </InputGroupAddon>
                </InputGroup>
                <Table striped>
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>User ID</th>
                        <th>Username</th>
                        <th>Location</th>
                        <th>Tweet</th>
                        <th>Score</th>
                        <th>Threat</th>
                        <th>Created</th>
                    </tr>
                    </thead>
                    <tbody>
                    {
                        tweets.map((t, index) => (
                            <tr key={index + 1}>
                                <td>{index}</td>
                                <td>{t?.userID}</td>
                                <td>{t?.username}</td>
                                <td>{t?.location}</td>
                                <td>{t?.text}</td>
                                <td>{t?.score}</td>
                                <td>{t?.threat}</td>
                                <td>{t?.created}</td>
                            </tr>
                        ))
                    }
                    </tbody>
                </Table>
            </div>
        );
    }
}

export default Tweets
