package zkutils

func (c *SdClient)Get(path string)  {
	c.conn.Get(path)
}
