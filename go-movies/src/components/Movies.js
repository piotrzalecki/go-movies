import React, { Component, Fragment } from 'react';
import { Link } from 'react-router-dom';

export default class Movies extends Component{

    state = {movies: [] };

    componentDidMount(){
        this.setState({
            movies: [
                {id:1, title: "Chlopaki nie placza", runtime: 110},
                {id:2, title: "Killer", runtime: 115},
                {id:3, title: "Poranek Kojota", runtime: 135},
                {id:4, title: "Rejs", runtime: 35},
            ]
        })
    }    

    render(){
        return(
            <Fragment>
                <h2>Choose a movie</h2>

                <ul>
                    {this.state.movies.map((m) => (
                        <li key={m.id}>
                            <Link to={`/movies/${m.id}`}>{m.title}</Link>
                        </li>

                    ))}
                </ul>
            </Fragment>
        );
    }
}