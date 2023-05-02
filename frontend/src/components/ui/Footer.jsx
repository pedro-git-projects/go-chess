const Footer = () => {
  const date = new Date()
  let year = date.getFullYear()
  console.log(year)
  return (
    <footer className="bg-[#00ADD8] py-4">
      <div className="container mx-auto text-center text-white font-bold">
        <p>
          © {year} Go Chess <br />
          Go get the source code at the{" "}
          <a
            className="text-rose-500"
            href="https://github.com/pedro-git-projects/projeto-integrado-frontend"
            target="_blank"
          >
            Github
          </a>
        </p>
      </div>
    </footer>
  )
}

export default Footer
