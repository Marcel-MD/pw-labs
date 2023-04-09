export type Quiz = {
	id: number;
	title: string;
	questions: Question[];
};

export type Question = {
	id: number;
	question: string;
	answers: string[];
};

export type QuizInfo = {
	id: number;
	title: string;
	questions_count: number;
	completed?: boolean;
	score?: number;
};

export type Answer = {
	id: number;
	correct_answer: string;
	correct: boolean;
};
