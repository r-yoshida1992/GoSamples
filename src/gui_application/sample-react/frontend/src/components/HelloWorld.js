import React, { useState } from 'react';
import Modal from 'react-modal';

function HelloWorld() {
	const [showModal, setShowModal] = useState(false);
	const [result, setResult] = useState(null);

	const handleOpenModal = () => {
		setShowModal(true);

		window.backend.basic().then((result) => setResult(result));
	};

	const handleCloseModal = () => {
		setShowModal(false);
	};

	const modalStyle = {
		overlay: {
			position: "fixed",
			top: 0,
			left: 0,
			backgroundColor: "rgba(0,0,0,0.85)"
		},
		content: {
			position: "absolute",
			top: "5rem",
			left: "5rem",
			right: "5rem",
			bottom: "5rem",
			backgroundColor: "paleturquoise",
			borderRadius: "1rem",
			padding: "1.5rem"
		}
	};

	return (
		<div className="App">
			<button onClick={() => handleOpenModal()} type="button">
				Hello
      </button>
			<Modal
				appElement={document.getElementById("app")}
				isOpen={showModal}
				contentLabel="Minimal Modal Example"
				style={modalStyle}
			>
				<p>{result}</p>
				<button onClick={() => handleCloseModal()}>Close Modal</button>
			</Modal>
		</div>
	);
}

export default HelloWorld;
