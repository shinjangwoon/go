import axios from 'axios';
import React from 'react';
import { Container, Row, Col } from "react-bootstrap";

import { useState, useEffect } from "react";
import { Link } from 'react-router-dom';
import Spinner from 'react-bootstrap/Spinner';

const Home = () => {
    const [apiData, setApiData] = useState(false); 
    const [loading, setLoading] = useState(true);
    useEffect(() => {
      const fetchData = async () => {
  
        try {
          const apiUrl = process.env.REACT_APP_API_ROOT;
          const response = await axios.get(apiUrl);
          
          if (response.status === 200) {
            if (response?.data.statusText === "OK") {
              setApiData(response?.data?.blog_records);
            }
          }
          setLoading(false);
        } catch (error) {
          console.log(error.response);
        }
  
      };
  
      fetchData();
      return () => {};
    }, []);
  
  
    console.log(apiData);

    if (loading) {
      return (
      <>
        <Container className='spinner'>
          <Spinner animation='grow' />
        </Container>
      </>)
    }

    return (
        <Container>
            <Row>
                <Col xs="12" className="py-2">
                    <h1 className="text-center">
                        React Application with Go fiber Backend
                    </h1>
                </Col>

                {apiData && 
                    apiData.map((record, index) => (
                    <Col key={index} xs="4" className="py-5 box">
                        <div className="title">
                            <Link to={`/blog/${record.id}`}>{record.title}</Link>
                            </div>
                        <div>{record.post}</div>
                    </Col>

                ))}
            </Row>
        </Container>
    );
};

export default Home