function Home() {    
    return (
        <div style={{
          fontSize: 20,
          textAlign: "center",
        }}>
             <h1>Welcome to MyUniverse!</h1>
            <h3>There you can rate books, movies, games and music, and share your profile with friends!</h3>
            <div style={{
                display: "inline-block",
                margin: 10,
            }}>
                <h3 className="link"><a href="signup">Sign Up</a></h3>
                <h3 className="link"><a href="login">Sign In</a></h3>
            </div>
            
        </div>
        
    );
}

export default Home;