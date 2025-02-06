import React, { useEffect, useState } from "react";
import { Table } from "antd";

const API_URL = "http://localhost:8080/pings";

interface PingData {
  ip: string;
  time: number;
  last_active: string;
}

const App: React.FC = () => {
  const [data, setData] = useState<PingData[]>([]);

  useEffect(() => {
    fetch(API_URL)
      .then((res) => res.json())
      .then((data) => setData(data));
  }, []);

  const columns = [
    { title: "IP Address", dataIndex: "ip", key: "ip" },
    { title: "Ping Time (ms)", dataIndex: "time", key: "time" },
    { title: "Last Active", dataIndex: "last_active", key: "last_active" },
  ];

  return <Table dataSource={data} columns={columns} rowKey="ip" />;
};

export default App;
