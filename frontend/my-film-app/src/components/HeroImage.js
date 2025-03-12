import React, { useEffect, useState } from 'react';
import './HeroImage.css';

const HeroImage = ({ imageSrc }) => {
  const [offsetY, setOffsetY] = useState(0);

  useEffect(() => {
    const handleScroll = () => {
      setOffsetY(window.pageYOffset);
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  return (
    <div
      className="hero-image-container"
      style={{
        transform: `translateY(${offsetY * 0.5}px)`,
        opacity: Math.max(1 - offsetY / 500, 0.5),
      }}
    >
      <img
        src={imageSrc}
        alt="大图展示"
        className="hero-image"
      />
    </div>
  );
};

export default HeroImage;
