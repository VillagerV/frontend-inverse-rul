import React, { useState, useEffect } from 'react';
import axios from 'axios';

const DataFetcher = () => {
  const [data, setData] = useState(null);
  const [url, setUrl] = useState('');

  useEffect(() => {
    const fetchUrlData = async () => {
      if (url) {
        try {
          const response = await axios.get(url);
          setData(response.data);
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      }
    };

    fetchUrlData();
  }, [url]);

  useEffect(() => {
    const savedUrl = localStorage.getItem('savedUrl');
    if (savedUrl) {
      setUrl(savedUrl);
    }
  }, []);

  const handleUrlChange = (e) => {
    const newUrl = e.target.value;
    setUrl(newUrl);
    localStorage.setItem('savedUrl', newUrl);
  };

  return (
    <div>
      <h1>Data Fetcher</h1>
      <label htmlFor="url-input">URL:</label>
      <input
        id="url-input"
        type="text"
        value={url}
        onChange={handleUrlChange}
        placeholder="Enter URL here"
      />
      {data && (
        <div>
          <h2>Results</h2>
          <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};

export default DataFetcher;
