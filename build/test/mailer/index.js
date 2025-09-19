const express = require('express');
const nodemailer = require('nodemailer');

const app = express();
app.use(express.json());

const transporter = nodemailer.createTransport({
  service: 'gmail',
  auth: {
    user: "oppenfamily24@gmail.com",
    pass: "gcbbwtwtfuhmxibg",
  },
});

app.get("/logo.png" , (req, res) => {
  res.sendFile(__dirname + '/QubeFlex.png');
});

app.post('/send', async (req, res) => {
  const { to, subject, matricula, contrasena } = req.body;

  if (!to || !subject || !matricula || !contrasena) {
    return res.status(400).json({ error: 'Faltan campos obligatorios.' });
  }

  const htmlContent = ` 
    <div style="font-family: system-ui; text-align: center; width: 300px; border-radius: 20px; background: #fff; color: #000; box-shadow: 0 4px 8px #F7F8FA; overflow: hidden; margin: 0 auto; border: 1px solid #fff;">
     <div style="height: 192px; width: 100%; display: flex; justify-content: center; background: linear-gradient(222deg, #fff 0%, #B4B8F3 100%); border-radius: 20px 20px 0 0;">
        <div style="display: flex;justify-content: center; align-items: center; align-text: center;background: rgba(247, 249, 250, 0);text-align: center;">
            <img src="cid:logo123" alt="Logo QubeFlex"  style="padding: 50px 10px 10px 10px ;width: 280px; height: 80px;">
        </div>
     </div>
      <div style="padding: 5px; text-align: center;">
        <h2 style="font-weight: 500; font-size: 22px;">Bienvenido a QubeFlex Droply</h2>
        <p>Tu cuenta ha sido dada de alta con éxito.</p>
        <p>Tus credenciales de acceso son:</p>
        <p style="font-weight: 700; font-size: 16px; color: #8780E2;">
          Matrícula: <span style="color: #000;">${matricula}</span><br>
          Contraseña: <span style="color: #000;">${contrasena}</span>
        </p>
      </div>
    </div>  
  `;

  try {
    await transporter.sendMail({
      from: "oppenfamily24@gmail.com",
      to,
      subject,
      html: htmlContent,
      attachments: [
        {
            filename: 'logo.png',
            path: './QubeFlex.png', // o un Buffer si viene de un archivo subido
            cid: 'logo123' // Este es el ID que usaste en src="cid:logo123"
        }
      ]
    });
    res.status(200).json({ message: 'Correo enviado exitosamente' });
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'No se pudo enviar el correo' });
  }
});


const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Servidor corriendo en el puerto ${PORT}`));
