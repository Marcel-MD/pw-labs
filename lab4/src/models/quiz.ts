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
};
