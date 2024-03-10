---
license: apache-2.0
language:
- en
---
<!DOCTYPE html>
<html lang="en">
  <head>
    <!-- MADE BY PotatoOff & LLM | https://huggingface.co/PotatoOff | Have fun and dont remove the credits <3 -->
    <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>HamSter v0.2</title>
      <link href="https://fonts.googleapis.com/css2?family=Quicksand:wght@400;500;600&display=swap" rel="stylesheet">
      <style>
          body {
              font-family: 'Quicksand', sans-serif;
              background-color: #1A202C;
              color: #F7FAFC;
              margin: 0;
              padding: 20px;
              font-size: 16px;
          }
          .container {
              width: 100%;
              margin: auto;
              background-color: #2D3748;
              padding: 20px;
              border-radius: 10px;
              box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
          }
          .header {
              display: flex;
              align-items: flex-start;
              gap: 20px;
          }
          .header h1 {
              font-size: 20px;
              color: #E2E8F0;
          }
          .header img {
              flex-shrink: 0;
              margin-left: 25%;
              width: 50%;
              max-width: 50%;
              border-radius: 15px;
              transition: filter 0.4s ease;
          }
          .header img:hover {
              filter: blur(2px); /* Apply a stronger blur on hover */
          }
          .info {
              flex-grow: 1;
              background-color: #2D3748;
              color: #CBD5E0;
              font-family: 'Fira Code', 'JetBrains Mono', monospace;
              padding: 15px;
              border-radius: 10px;
              box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
              font-size: 14px;
              line-height: 1.7;
              overflow-x: auto;
              margin-top: 40px;
              border: 2px solid #4A90E2;
              transition: box-shadow 0.3s ease;
              position: relative; /* Ensure proper stacking */
          }
          .info:hover {
              box-shadow: 0 4px 13px rgba(0, 0, 0, 0.6), 0 0 24px rgba(74, 144, 226, 0.6);
          }
          .info-img {
              width: 100%; /* Adjust width as per your layout needs */
              max-width: 400px; /* Max width to ensure it doesn't get too large */
              max-height: 100%; /* Adjust height proportionally */
              border-radius: 10px;
              box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
              margin-left: 5%; /* Align to the right */
              margin-right: 0%; /* Keep some space from the text */
              display: block; /* Ensure it's properly block level for margins to work */
              float: right; /* Keep it to the right */
          }
        .button {
            display: inline-block;
            background-image: linear-gradient(145deg, #F96167 0%, #F0F2D7 100%);
            color: #F0F0F0;
            padding: 16px 24px; /* Increased padding for bigger buttons */
            border: none;
            border-radius: 10px;
            cursor: pointer;
            text-decoration: none;
            margin-left: 7%;
            transition: transform 0.3s ease, box-shadow 0.3s ease, background-image 0.3s ease, color 0.3s ease, border-radius 0.3s ease; /* Enhanced transitions */
            font-weight: bold; /* Make the text bold */
            box-shadow: 0 2px 15px rgba(0, 0, 0, 0.2); /* Subtle shadow for depth */
        }
        .button:hover {
            background-image: linear-gradient(145deg, #FB1A3E 0%, #F38555 100%); /* Vibrant to light pink gradient */
            transform: scale(1.1); /* Increase size for more emphasis */
            box-shadow: 0 10px 30px rgba(249, 97, 103, 0.8); /* More pronounced glowing effect */
            color: #FFFFFF; /* Brighten the text color slightly */
            border-radius: 15px; /* Soften the corners a bit more for a pill-like effect */
        }
        @keyframes pulse {
          0% {
              transform: scale(1);
              opacity: 1;
          }
          50% {
              transform: scale(1.05);
              opacity: 0.85;
          }
          100% {
              transform: scale(1);
              opacity: 1;
          }
      }
      </style>
  </head>

  <body>
    <div class="container">
      <div class="header">
        <div class="info" style="margin-top: 5px;">
          <img src="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/RPNN3Hs6mMXe25viwZklX.png" alt="Image">
          <h1 class="product-name" style="margin: 10px">Meet HamSter-0.1 🐹</h1>
          <p>
            👋 Uncensored fine tune model roleplay focused of "mistralai/Mistral-7B-v0.2" and first model of the HamSter series. Made with the help of my team <a href="https://huggingface.co/ConvexAI" target="_blank">ConvexAI.</a><br><br>
            🚀 For optimal performance, I recommend using a detailed character card! (There is NSFW chub.ai) Check out <a href="https://chub.ai" target="_blank">Chub.ai</a> for some character cards.<br><br>
            🤩 Uses the Llama2 prompt template with chat instructions.<br><br>
            🔥 Produce spicy content.<br><br>
            😄 -> Check out <a href="https://huggingface.co/PotatoOff/HamSter-0.2" target="_blank">HamSter 0.2</a> latest model of the HamSter series. Check it out!<br>
          </p>
          <div>
            <a href="https://huggingface.co/collections/PotatoOff/hamster-01-65a31043b7897304be56474d" class="button">HamSter 0.1 Quants</a>
            <a href="https://discord.com/invite/9y7KxZxcZx" class="button">Discord Server</a>
          </div>
        </div>
    </div>
    <div style="overflow: hidden; position: relative">
      <div class="info"style="overflow: hidden; margin:-left 0% margin-top: 20px;">       
        <a href="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/Mxxwp63AHWhxfSuC7ljsB.png" target="_blank">
          <img src="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/Mxxwp63AHWhxfSuC7ljsB.png" alt="Roleplay Test" style="width: auto; max-width: 37%; max-height: 100%; border-radius: 10px; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2); margin-left: 0%; display: block; float: right;">
        </a>
        <h2 style="margin-top: 0;">I had good results with these parameters:</h2>
        <ul style="margin-top: 0;">
          <p>> temperature: 0.8 <</p>
          <p>> top_p: 0.75</p>
          <p>> min_p: 0</p>
          <p>> top_k: 0</p>
          <p>> repetition_penalty: 1.05</p>
        </ul>
      </div>
    </div>
    <div style="overflow: hidden; position: relative;">
      <div class="info" style="overflow: hidden; margin-top: 20px;">
        <h2 style="margin-top: 0;">BenchMarks on OpenLLM Leaderboard</h2>
        <a href="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/ZDFGednAQjtPQDjyvmlRU.webp" target="_blank">
          <img src="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/ZDFGednAQjtPQDjyvmlRU.webp" alt="OPEN LLM BENCHMARK" style="info-img; border-radius: 10px">
        </a>
        <p>More details: <a href="https://huggingface.co/datasets/open-llm-leaderboard/details_PotatoOff__HamSter-0.1" target="_blank">HamSter-0.1 OpenLLM BenchMarks</a></p>
      </div>
    </div>
    <div style="overflow: hidden; position: relative;">
      <div class="info" style="overflow: hidden; margin-top: 20px;">
        <h2 style="margin-top: 0;">BenchMarks on Ayumi's LLM Role Play & ERP Ranking</h2>
        <a href="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/1z6u-_Iu3dXoAo-ia0KKl.png" target="_blank">
          <img src="https://cdn-uploads.huggingface.co/production/uploads/64e7616c7df33432812e3c57/1z6u-_Iu3dXoAo-ia0KKl.png" alt="Ayumi's LLM Role Play & ERP Ranking" class="info-img" style="width: 100%; height: auto;">
        </a>
        <p>More details: <a href="http://ayumi.m8geil.de/results_v3/model_resp_DL_20240113_7B-Q6_K_HamSter_0_1.html">Ayumi's LLM RolePlay & ERP Rankin HamSter-0.1 GGUF version Q6_K</a></p>
      </div>
    </div>
      <div style="font-family: 'Arial', sans-serif; font-weight: bold; text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.5);">
        <p style="display: inline; font-size: 17px; margin: 0;">Have Fun</p>
        <p style="display: inline; color: #E2E8F0; margin-bottom: 20px; animation: pulse 2s infinite; font-size: 17px;">💖</p>
      </div>
    </div>
  </body>
</html>
